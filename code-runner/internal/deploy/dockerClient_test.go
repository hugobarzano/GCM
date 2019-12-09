package deploy

import (
	"code-runner/internal/models"
	"fmt"
	"testing"
	"time"
)

func TestPull(t *testing.T) {

	//ctx:=context.Background()
	app:= &models.App{
		Name:"cat",
		Owner:"hugobarzano",
	}

	dockerApp:= DockerApp{
		App:app,
	}



	dockerApp.Start("")

	time.Sleep(time.Second*5)
	err:=dockerApp.ContainerStop()
	if err!=nil{
		fmt.Printf(err.Error())
	}
}

//func TestPull(t *testing.T) {
//	ctx := context.Background()
//	//cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
//
//
//	cli, err :=client.NewEnvClient()
//	authConfig := types.AuthConfig{
//		Username: "gg",
//		RegistryToken: "",
//	}
//
//	body,err:=cli.RegistryLogin(ctx,authConfig)
//	fmt.Println(body)
//
//	authConfig.IdentityToken=body.IdentityToken
//	encodedJSON, err := json.Marshal(authConfig)
//	if err != nil {
//		panic(err)
//	}
//	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
//
//	out, err := cli.ImagePull(ctx, "docker.pkg.github.com/hugobarzano/cat/hugobarzano.cat:latest", types.ImagePullOptions{RegistryAuth: authStr})
//	if err != nil {
//		panic(err)
//	}
//
//	defer out.Close()
//	io.Copy(os.Stdout, out)
//}
//
//
//func TestClient_Pull(t *testing.T) {
//
//	deployClient := GetDockerClient()
//
//	opt:=types.ImagePullOptions{
//		RegistryAuth:"7040fba5670af31299fe2c448778df8aa59fad69",
//
//	}
//
//	ctx:=context.Background()
//
//	auth:= types.AuthConfig{
//		Username: "hugobarzano",
//		RegistryToken: "7040fba5670af31299fe2c448778df8aa59fad69",
//	}
//
//	body,err:=deployClient.docker.RegistryLogin(ctx,auth)
//	fmt.Println(body)
//
//
//	resp, err := deployClient.docker.ImagePull(ctx, "docker.pkg.github.com/hugobarzano/cat/hugobarzano.cat:latest", opt)
//
//
//	if err != nil{
//		fmt.Println(err)
//	}
//
//	buf := new(bytes.Buffer)
//	buf.ReadFrom(resp)
//	newStr := buf.String()
//
//	fmt.Printf(newStr)
//
//	fmt.Println(resp)
//
//}
//
//func TestClient_RunContainer(t *testing.T) {
//	//deployClient := GetDockerClient()
//	//deployClient.RunContainerFromImage(context.Background(),"alpine")
//	//
//	//d:=`FROM    httpd:2.4
//	//			MAINTAINER    hugobarzano `
//	//app := &models.App{
//	//	Name:  "apache",
//	//	Owner: "hugobarzano",
//	//}
//
//	//err:=deployClient.BuildImage(context.Background(),*app)
//
//	//fmt.Println(err)
//}
//
//func Test_New(t *testing.T)  {
//
//	app := &models.App{
//		Name:  "apache",
//		Owner: "hugobarzano",
//	}
//
//	deployClient := GetDockerClient()
//
//	repoPath:=tools.GetAppLocalPath(*app)
//	dockerBuildContext, err := os.Open(repoPath)
//	defer dockerBuildContext.Close()
//	if err!=nil{
//		fmt.Println(err)
//	}
//
//	opt := types.ImageBuildOptions{}
//	response, err := deployClient.docker.ImageBuild(context.Background(), dockerBuildContext, opt)
//	if err == nil {
//		fmt.Printf("Error building, %v", err)
//	}
//	fmt.Println("response")
//	fmt.Println(response)
//}
