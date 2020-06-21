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
	READY    AppStatus = "ready"
	RUNNING  AppStatus = "running"
	STOPPED  AppStatus = "stopped"
	DELETING AppStatus = "deleting"
)

type App struct {
	Name       string            `bson:"_id"  json:"name"`
	Repository string            `bson:"repo" json:"repo"`
	Spec       map[string]string `bson:"spec" json:"spec"`
	Des        string            `bson:"des" json:"des,omitempty"`
	Url        string            `bson:"url"  json:"url"`
	Owner      string            `bson:"owner"  json:"owner"`
	Status     AppStatus         `bson:"status"  json:"status"`
	Img        string            `bson:"-" json:"img"`
}

func (app *App) IsRunning() bool {
	return app.Status == RUNNING
}

func (app *App) GetImageName(version string) string {
	imageName := fmt.Sprintf("%v.%v:%v", app.Owner, app.Name, version)
	return imageName
}

func (app *App) GetPKGName(version string) string {
	pkg := app.GetImageName(version)
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
	case constants.SinglePage:
		appUrl = fmt.Sprintf("http://%v:%v",
			config.GetConfig().DeployAddress, port)
	case constants.ApiRest:
		switch app.Spec["tech"] {
		case "js":
			appUrl = fmt.Sprintf("http://%v:%v/api/home",
				config.GetConfig().DeployAddress, port)
		default:
			appUrl = fmt.Sprintf("http://%v:%v/api",
				config.GetConfig().DeployAddress, port)
		}
	case constants.DataService:
		switch app.Spec["tech"] {
		case "mongodb":
			appUrl = fmt.Sprintf("mongo mongodb://%v:%v",
				config.GetConfig().DeployAddress, port)
		case "mysql":
			appUrl = fmt.Sprintf("mysql -h %v --port %v -u root",
				config.GetConfig().DeployAddress, port)
		case "redis":
			appUrl = fmt.Sprintf("redis-cli -h %v -p %v",
				config.GetConfig().DeployAddress, port)
		}
	case constants.DevOpsService:
		appUrl = fmt.Sprintf("http://%v:%v",
			config.GetConfig().DeployAddress, port)
	}
	app.Url = appUrl
}
