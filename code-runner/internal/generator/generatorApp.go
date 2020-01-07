package generator

import (
	"code-runner/internal/deploy"
	"code-runner/internal/models"
	"code-runner/internal/store"
	"context"
	"fmt"
	googleGithub "github.com/google/go-github/github"
)

type GenApp struct {
	App               *models.App
	Github 			*googleGithub.Client
	Readme []byte
	Dockerfile []byte
	CI map[string][]byte
	Data map[string][]byte
	Local map[string][]byte
}

func (app *GenApp) InitializeCode(user string, token string) {

	ctx:=context.Background()
	app.InitGit(ctx,token)
	app.generateReadme()
	if err:=app.pushReadme(ctx,user); err!=nil{
		fmt.Printf("PushFile Error: %s", err.Error())
	}
	app.generateSourceCode()
	if err:=app.pushSourceCode(ctx,user); err!=nil{
		fmt.Printf("PushFile Error: %s", err.Error())
	}
	app.generateDockerfile()
	if err:=app.pushDockerfile(ctx,user);err!=nil{
		fmt.Printf("PushFile Error: %s", err.Error())
	}

	app.generateLocalUtils()
	if err:=app.pushLocalUtilsCode(ctx,user);err!=nil{
		fmt.Printf("PushFile Error: %s", err.Error())
	}

	app.generateCI()
	if err:=app.pushCI(ctx,user);err!=nil{
		fmt.Printf("PushFile Error: %s", err.Error())
	}

	dao:=store.InitMongoStore(ctx)
	app.App.Status=models.BUILDING
	_,err:=dao.UpdateApp(ctx,app.App)
	if err !=nil{
		fmt.Printf("DB Error: %s", err.Error())
	}

	dockerApp:=deploy.DockerApp{
		App:app.App,
	}
	go dockerApp.Start(token)
}

