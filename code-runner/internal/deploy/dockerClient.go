package deploy

import (
	"code-runner/internal/constants"
	"code-runner/internal/models"
	"code-runner/internal/store"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"golang.org/x/net/context"
	"io"
	"net"
	"os"
	"strconv"
	"time"
)

type DockerApp struct {
	App               *models.App
	Client            *client.Client
	AuthConfig        types.AuthConfig
	AuthConfigEncoded string
	Config            container.ContainerCreateCreatedBody

}

func registryAuthentication(user,password string) types.RequestPrivilegeFunc {
	return func() (string, error) {
		authConfig := types.AuthConfig{
			Username: user,
			Password: password,
			ServerAddress: constants.DockerRegistry,
		}
		buf, err := json.Marshal(authConfig)
		if err != nil {
			return "", err
		}
		base64config:=base64.URLEncoding.EncodeToString(buf)
		if err != nil {
			return "", err
		}
		return base64config, nil
	}
}

func (appDocker *DockerApp) PrepareRegistry(ctx context.Context,password string) error {
	appDocker.AuthConfig = types.AuthConfig{
		Username: appDocker.App.Owner,
		Password: password,
		ServerAddress: constants.DockerRegistry,
	}
	resp, err := appDocker.Client.RegistryLogin(ctx, appDocker.AuthConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("Status:\t", resp.Status)
	if resp.IdentityToken != "" {
		appDocker.AuthConfig.IdentityToken = resp.IdentityToken
	}
	buf, err := json.Marshal(appDocker.AuthConfig)
	appDocker.AuthConfigEncoded=base64.URLEncoding.EncodeToString(buf)
	return err
}

func (appDocker *DockerApp) ImagePull(ctx context.Context,token string) error {

	opts := types.ImagePullOptions{
		RegistryAuth: appDocker.AuthConfigEncoded,
		PrivilegeFunc: registryAuthentication(appDocker.App.Owner,token),
	}

	pkgAddr:=appDocker.App.GetPKGName()

	img, err := appDocker.Client.ImagePull(ctx,pkgAddr, opts)
	//var err error
	//var img io.ReadCloser
	timeout:=time.Duration(5*time.Minute)
	start := time.Now()
	for img==nil  {
		fmt.Printf("\nimage not ready: %v",appDocker.App.Status)
		img, err = appDocker.Client.ImagePull(ctx,pkgAddr, opts)
		if time.Since(start)>=timeout{
			return errors.New("Error pulling image after "+timeout.String())
		}

	}
	_,_=io.Copy(os.Stdout, img)
	if err != nil {
		return err
	}
	dao:=store.InitMongoStore(ctx)
	appDocker.App.Status=models.READY
	_,err=dao.UpdateApp(ctx,appDocker.App)
	if err !=nil{
		fmt.Printf("DB Error: %s", err.Error())
		return err
	}
	return nil
}


func (appDocker *DockerApp) ContainerStart() error {
	fmt.Printf("config")
	fmt.Println(appDocker.Config)
	return appDocker.Client.ContainerStart(context.Background(), appDocker.Config.ID, types.ContainerStartOptions{})
}

func (appDocker *DockerApp) ContainerStop() error {

	timeOut:=0 * time.Second
	fmt.Println("DOCKER ID: "+appDocker.App.Spec["dockerId"])
	return 	appDocker.Client.ContainerStop(context.Background(),appDocker.App.Spec["dockerId"],&timeOut)
}

func (app *DockerApp) Start(token string) {

	var err error
	err = app.Initialize()
	if err != nil {
		fmt.Println("Initialize error: "+err.Error())
	}

	err = app.PrepareRegistry(context.Background(),token)
	if err != nil {
		fmt.Println("PrepareRegistry error: "+err.Error())
	}

	err = app.ImagePull(context.Background(),token)
	if err != nil {
		fmt.Println("ImagePull error: "+err.Error())
	}

	err = app.ContainerCreate(context.Background())
	if err != nil {
		fmt.Println("ContainerCreate error: "+err.Error())
	}
}


func getAvailablePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}



func (app *DockerApp) Initialize() error {
	var err error
	app.Client, err = client.NewEnvClient()
	return err
}


func (appDocker *DockerApp) ContainerCreate(ctx context.Context) error {
	availablePort,_:=getAvailablePort()
	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: strconv.Itoa(availablePort),
	}
	containerPort, err := nat.NewPort("tcp", appDocker.App.Spec["port"])
	if err != nil {
		panic("Unable to get the port")
	}
	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}


	containerObj, err := appDocker.Client.ContainerCreate(context.Background(),
		&container.Config{ Image: appDocker.App.GetPKGName()},
		&container.HostConfig{
			PortBindings: portBinding},
		nil, "")

	if err != nil {
		//panic("Unable to Create container")
		return err
	}

	err = appDocker.Client.ContainerStart(ctx,
		containerObj.ID,
		types.ContainerStartOptions{});

	if err != nil {
		//panic("Unable to Start container")
		return err
	}
	appDocker.Config=containerObj
	appDocker.App.Spec["dockerId"]=containerObj.ID
	appDocker.App.Status=models.RUNNING
	appDocker.App.SetDeployURL(strconv.Itoa(availablePort))

	dao:=store.InitMongoStore(ctx)
	_,err=dao.UpdateApp(ctx,appDocker.App)
	if err != nil {
		//panic("Unable to update DB container")
		return err
	}
	return err
}
