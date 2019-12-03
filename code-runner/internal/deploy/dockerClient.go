package deploy

import (
	"code-runner/internal/models"
	"code-runner/internal/tools"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"golang.org/x/net/context"
	"io"
	"os"
)

type Client struct {
	docker *client.Client
}

func GetDockerClient() *Client {
	cli, err :=client.NewEnvClient()
	if err != nil {
		fmt.Printf("Error creating docker client: %s",err.Error())
		return nil
	}
	dockerClient:=&Client{
		docker:cli,
	}
	return dockerClient
}

func (c *Client)BuildImage(ctx context.Context,app models.App,sha string) error{

	repoPath:=tools.GetAppLocalPath(app)
	fmt.Println(repoPath)
	dockerBuildContext, err := os.Open(repoPath)
	defer dockerBuildContext.Close()
	if err!=nil{
		return err
	}
	dockerfilePath:=fmt.Sprintf("%v-%v-%v/Dockerfile",app.Owner,app.Name,sha)

	opt := types.ImageBuildOptions{
		Dockerfile:   dockerfilePath,
		Tags: []string{app.Name},
		SuppressOutput: false,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
		NoCache:        true,
	}

	response, err := c.docker.ImageBuild(ctx,
		dockerBuildContext, opt)
	if err != nil {
		fmt.Printf("Error building, %v", err)
		return err
	}
	fmt.Println("response")
	fmt.Println(response)
	_,err=io.Copy(os.Stdout, response.Body)
	//if err!=nil{
	//	fmt.Println("Error on copy stdout")
	//	return err
	//}
	return nil
}

func (c *Client)RunContainerFromImage(ctx context.Context,image string)  {
	imageName := "docker.io/library/"+image

	out, err := c.docker.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		fmt.Println("Error pulling imagen")
		fmt.Println(err)
	}
	_,err=io.Copy(os.Stdout, out)
	if err!=nil{
		fmt.Println("Error on copy stdout")
	}

	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: "8001",
	}

	containerPort, err := nat.NewPort("tcp", "80")
	if err != nil {
		panic("Unable to get the port")
	}
	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}


	containerObj, err := c.docker.ContainerCreate(ctx,
		&container.Config{ Image: image},
		&container.HostConfig{
			PortBindings: portBinding},
			nil, "")
	if err != nil {
		panic(err)
	}

	if err := c.docker.ContainerStart(ctx, containerObj.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(containerObj.ID)
	
}
