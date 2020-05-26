package generator

import (
	"code-runner/internal/constants"
	"code-runner/internal/models"
	"code-runner/internal/store"
	"context"
	"fmt"
	googleGithub "github.com/google/go-github/v31/github"
	"log"
)

type GenApp struct {
	App        *models.App
	Github     *googleGithub.Client
	Readme     []byte
	License    []byte
	Dockerfile []byte
	CI         map[string][]byte
	Data       map[string][]byte
	Local      map[string][]byte
}

func (app *GenApp) InitializeCode(user string, token string, mail string) {

	ctx := context.Background()
	app.InitGit(ctx, token)
	app.generateReadme()
	if err := app.pushReadme(ctx, user, mail); err != nil {
		log.Println("PushFile Error: %s", err.Error())
	}

	app.generateLicense()
	if err := app.pushLicense(ctx, user, mail); err != nil {
		log.Println("PushFile Error: %s", err.Error())
	}

	app.generateSourceCode()
	if err := app.pushSourceCode(ctx, user, mail); err != nil {
		log.Println("PushFile Error: %s", err.Error())
	}
	app.generateDockerfile()
	if err := app.pushDockerfile(ctx, user, mail); err != nil {
		log.Println("PushFile Error: %s", err.Error())
	}

	app.generateLocalTools()
	if err := app.pushLocalTools(ctx, user, mail); err != nil {
		log.Println("PushFile Error: %s", err.Error())
	}

	app.generateCI()
	if err := app.PushCI(ctx, user, mail); err != nil {
		log.Println(fmt.Sprintf("PushFile Error: %s", err.Error()))
	}

	app.App.Status = models.BUILDING
	_, err := store.ClientStore.UpdateApp(ctx, app.App)
	if err != nil {
		log.Println(fmt.Sprintf("DB Error: %s", err.Error()))
	}

	dockerApp := DockerApp{
		App: app.App,
	}
	go dockerApp.ContainerStart(token)
}

func (app *GenApp) ReGenerateCode(user string, token string, mail string) {

	ctx := context.Background()
	app.InitGit(ctx, token)
	app.generateReadme()
	if err := app.pushReadme(ctx, user, mail); err != nil {
		log.Println("PushFile Error: %s", err.Error())
	}

	app.generateLicense()
	if err := app.pushLicense(ctx, user, mail); err != nil {
		log.Println("PushFile Error: %s", err.Error())
	}

	app.generateSourceCode()
	if err := app.pushSourceCode(ctx, user, mail); err != nil {
		log.Println("PushFile Error: %s", err.Error())
	}
	app.generateDockerfile()
	if err := app.pushDockerfile(ctx, user, mail); err != nil {
		log.Println("PushFile Error: %s", err.Error())
	}

	app.generateLocalTools()
	if err := app.pushLocalTools(ctx, user, mail); err != nil {
		log.Println("PushFile Error: %s", err.Error())
	}

	app.generateCI()
	if err := app.PushCI(ctx, user, mail); err != nil {
		log.Println(fmt.Sprintf("PushFile Error: %s", err.Error()))
	}

	app.App.Status = models.BUILDING
	_, err := store.ClientStore.UpdateApp(ctx, app.App)
	if err != nil {
		log.Println(fmt.Sprintf("DB Error: %s", err.Error()))
	}

	dockerApp := DockerApp{
		App: app.App,
	}
	go dockerApp.ContainerRegenerate(token)
}

func (app *GenApp) generateSourceCode() {

	switch tech := app.App.Spec["tech"]; tech {
	case "apacheStatic":
		app.generateApacheSinglePageCode()
	case "nodeStatic":
		app.generateNodeSinglePageCode()
	case "mongodb":
		app.generateMongoService()
	case "mysql":
		app.generateMysqlService()
	case "redis":
		app.generateRedisService()
	case "jenkins":
		app.generateJenkinsService()
	default:
		log.Println("NOT SUPPORTED")
	}

	switch app.App.Spec["nature"] {
	case constants.ApiRest:
		app.generateApiService()
	default:
		log.Println("NOT SUPPORTED")
	}
}

func (app *GenApp) pushSourceCode(ctx context.Context, user, mail string) error {

	var commitMsg string
	var fileOptions *googleGithub.RepositoryContentFileOptions
	var err error

	for file, content := range app.Data {
		commitMsg = "Generating " + file
		fileOptions = BuildFileOptions(commitMsg, user, mail, content)
		_, err = app.CommitFile(ctx, file, fileOptions)
		if err != nil {
			return err
		}
	}
	return nil
}
