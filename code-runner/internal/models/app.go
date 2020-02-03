package models

import (
	"code-runner/internal/config"
	"code-runner/internal/constants"
	"fmt"
	"os"
)

type AppStatus string

const (
	INIT     AppStatus = "generating"
	BUILDING AppStatus = "building"
	READY    AppStatus = "readytorun"
	RUNNING  AppStatus = "running"
	STOPPED  AppStatus = "stopped"
)

type App struct {
	Name       string    `bson:"_id"  json:"name"`
	Repository string    `bson:"repo" json:"repo"`
	Spec       map[string]string    `bson:"spec" json:"spec"`
	Des        string    `bson:"des" json:"des,omitempty"`
	Url        string    `bson:"url"  json:"url"`
	Owner      string    `bson:"owner"  json:"owner"`
	Status     AppStatus `bson:"status"  json:"status"`
}

func (app *App) IsRunning() bool {
	return app.Status==RUNNING
}

func (app *App) GetImageName() string {
	imageName := fmt.Sprintf("%v.%v:latest", app.Owner, app.Name)
	return imageName
}

func (app *App) GetPKGName() string {
	pkg := app.GetImageName()
	pkgAddr := fmt.Sprintf("%v/%v/%v/%v", constants.DockerRegistry, app.Owner, app.Name, pkg)
	return pkgAddr
}

func (app *App) GetLocalPath() string {
	return fmt.Sprintf("%v%v/%v/%v.tar.gz",
		os.TempDir(), app.Owner, app.Name, app.Name)
}

func (app *App) SetDeployURL(port string) {

	var appUrl string
	switch app.Spec["nature"] {
	case "staticApp":
		appUrl=fmt.Sprintf("http://%v:%v",
			config.GetConfig().DeployAddress, port)
	case "mongodb":
		appUrl=fmt.Sprintf("mongodb://%v:%v",
			config.GetConfig().DeployAddress, port)
	default:
		appUrl=fmt.Sprintf("%v:%v",
			config.GetConfig().DeployAddress, port)
	}
	app.Url=appUrl
}