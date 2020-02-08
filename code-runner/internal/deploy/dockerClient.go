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
	"log"
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

func registryAuthentication(user, password string) types.RequestPrivilegeFunc {
	return func() (string, error) {
		authConfig := types.AuthConfig{
			Username:      user,
			Password:      password,
			ServerAddress: constants.DockerRegistry,
		}
		buf, err := json.Marshal(authConfig)
		if err != nil {
			return "", err
		}
		base64config := base64.URLEncoding.EncodeToString(buf)
		if err != nil {
			return "", err
		}
		return base64config, nil
	}
}

func (appDocker *DockerApp) prepareRegistry(ctx context.Context, password string) error {
	appDocker.AuthConfig = types.AuthConfig{
		Username:      appDocker.App.Owner,
		Password:      password,
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
	appDocker.AuthConfigEncoded = base64.URLEncoding.EncodeToString(buf)
	return err
}

func (appDocker *DockerApp) imagePull(ctx context.Context, token string) error {

	opts := types.ImagePullOptions{
		RegistryAuth:  appDocker.AuthConfigEncoded,
		PrivilegeFunc: registryAuthentication(appDocker.App.Owner, token),
	}

	pkgAddr := appDocker.App.GetPKGName()

	img, err := appDocker.Client.ImagePull(ctx, pkgAddr, opts)
	timeout := time.Duration(5 * time.Minute)
	start := time.Now()
	for img == nil {
		fmt.Printf("image: %v not ready. Status: %v\n", appDocker.App.GetPKGName(),appDocker.App.Status)
		img, err = appDocker.Client.ImagePull(ctx, pkgAddr, opts)
		time.Sleep(time.Second*5)
		if time.Since(start) >= timeout {
			return errors.New("Error pulling image after " + timeout.String())
		}

	}
	_, _ = io.Copy(os.Stdout, img)
	if err != nil {
		return err
	}
	dao := store.InitMongoStore(ctx)
	appDocker.App.Status = models.READY
	_, err = dao.UpdateApp(ctx, appDocker.App)
	if err != nil {
		fmt.Printf("DB Error: %s", err.Error())
		return err
	}
	return nil
}


func (appDocker *DockerApp) ContainerStop(ctx context.Context) error {

	timeOut := 0 * time.Second
	fmt.Println("DOCKER ID: " + appDocker.App.Spec["dockerId"])
	err:=appDocker.Client.ContainerStop(ctx, appDocker.App.Spec["dockerId"], &timeOut)
	return err
}

func (appDocker *DockerApp) ContainerRemove(ctx context.Context) error {

	err:=appDocker.Client.ContainerRemove(ctx,appDocker.App.Spec["dockerId"],types.ContainerRemoveOptions{
		Force:true,
		RemoveLinks:false,
		RemoveVolumes:true,
	})
	return err
}

func (appDocker *DockerApp) ImageRemove(ctx context.Context) error {

	del,err:=appDocker.Client.ImageRemove(ctx,appDocker.App.GetPKGName(),types.ImageRemoveOptions{
		Force:true,
	})

	fmt.Print(del)
	fmt.Printf(err.Error())
	return err
}

func (app *DockerApp) ContainerStart(token string) {

	var err error
	err = app.Initialize()
	if err != nil {
		fmt.Println("Initialize error: " + err.Error())
	}

	err = app.prepareRegistry(context.Background(), token)
	if err != nil {
		fmt.Println("prepareRegistry error: " + err.Error())
	}

	err = app.imagePull(context.Background(), token)
	if err != nil {
		fmt.Println("imagePull error: " + err.Error())
	}

	err = app.containerCreate(context.Background())
	if err != nil {
		fmt.Println("containerCreate error: " + err.Error())
	}
}

func getAvailablePort() string {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return "0"
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return "0"
	}
	defer l.Close()
	return strconv.Itoa(l.Addr().(*net.TCPAddr).Port)
}

func (app *DockerApp) Initialize() error {
	var err error
	app.Client, err = client.NewEnvClient()
	return err
}

func (appDocker *DockerApp) getPortBinding(tcpPort string) (nat.PortMap, error) {
	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: tcpPort,
	}
	containerPort, err := nat.NewPort("tcp", appDocker.App.Spec["port"])
	if err != nil {
		fmt.Print("Unable to get Container port: " + err.Error())
		return nil, err
	}

	portBinding := nat.PortMap{
		containerPort: []nat.PortBinding{
			hostBinding,
		},
	}
	return portBinding, nil

}

func (appDocker *DockerApp) containerCreate(ctx context.Context) error {

	ctx=context.Background()
	availablePort:=appDocker.App.Spec["port"]
	portBinding, err := appDocker.getPortBinding(availablePort)
	if err != nil {
		return err
	}

	containerObj, err := appDocker.Client.ContainerCreate(ctx,
		&container.Config{Image: appDocker.App.GetPKGName()},
		&container.HostConfig{
			PortBindings: portBinding},
		nil, "")

	if err != nil {
		return err
	}

	err = appDocker.Client.ContainerStart(ctx,
		containerObj.ID,
		types.ContainerStartOptions{})

	if err != nil {
		fmt.Println("Unable to start with client port")
		err:=appDocker.Client.ContainerRemove(ctx,containerObj.ID,types.ContainerRemoveOptions{
			Force:true,
		})
		if err != nil {
			return err
		}

		availablePort = getAvailablePort()
		portBinding, err = appDocker.getPortBinding(availablePort)
		if err != nil {
			return err
		}

		containerObj, err = appDocker.Client.ContainerCreate(ctx,
			&container.Config{Image: appDocker.App.GetPKGName()},
			&container.HostConfig{
				PortBindings: portBinding},
			nil, "")

		err = appDocker.Client.ContainerStart(ctx,
			containerObj.ID,
			types.ContainerStartOptions{})

		if err != nil {
			fmt.Println("Unable to start with available port")
			return err
		}
	}

	appDocker.Config = containerObj
	appDocker.App.Spec["dockerId"] = containerObj.ID
	appDocker.App.Status = models.RUNNING
	appDocker.App.SetDeployURL(availablePort)
	dao := store.InitMongoStore(ctx)
	_, err = dao.UpdateApp(ctx, appDocker.App)

	return err
}

func (appDocker *DockerApp) GetContainerLogReader(ctx context.Context) io.ReadCloser {
	reader, err := appDocker.Client.ContainerLogs(ctx,
		appDocker.App.Spec["dockerId"], types.ContainerLogsOptions{
			ShowStdout: true,
			ShowStderr: true,
			Details:    false,
			Follow:     true,
			Timestamps: false,
			Tail:       "1",
		})
	if err != nil {
		log.Print(err.Error())
	}
	return reader
}

func (appDocker *DockerApp) GetContainerLogById2(ctx context.Context, id string) io.ReadCloser {
	//appDocker.Client.ContainerInspect()
	reader, err := appDocker.Client.ContainerLogs(ctx, id,
		types.ContainerLogsOptions{
			ShowStdout: true,
			ShowStderr: true,
			Details:    true,
			Follow:     true,
			Timestamps: true,
			Tail:       "1",
		})
	if err != nil {
		log.Fatal(err)
	}
	return reader
}
