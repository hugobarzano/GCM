package deploy

import (
	"code-runner/internal/models"
	"code-runner/internal/tools"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"os"
	"testing"
)

func TestClient_RunContainer(t *testing.T) {
	deployClient := GetDockerClient()
	//deployClient.RunContainerFromImage(context.Background(),"alpine")
	//
	//d:=`FROM    httpd:2.4
	//			MAINTAINER    hugobarzano `
	app := &models.App{
		Name:  "apache",
		Owner: "hugobarzano",
	}

	err:=deployClient.BuildImage(context.Background(),app)

	fmt.Println(err)
}

func Test_New(t *testing.T)  {

	app := &models.App{
		Name:  "apache",
		Owner: "hugobarzano",
	}

	deployClient := GetDockerClient()

	repoPath:=tools.GetAppLocalPath(app)
	dockerBuildContext, err := os.Open(repoPath)
	defer dockerBuildContext.Close()
	if err!=nil{
		fmt.Println(err)
	}

	opt := types.ImageBuildOptions{}
	response, err := deployClient.docker.ImageBuild(context.Background(), dockerBuildContext, opt)
	if err == nil {
		fmt.Printf("Error building, %v", err)
	}
	fmt.Println("response")
	fmt.Println(response)
}
