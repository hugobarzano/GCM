package models

import (
	"code-runner/internal/constants"
	"fmt"
)

type App struct {
	Name       string `bson:"_id"  json:"name"`
	Repository string `bson:"repo" json:"repo"`
	Spec       string `bson:"spec" json:"spec"`
	Des        string `bson:"des" json:"des,omitempty"`
	Url        string `bson:"url"  json:"url"`
	Owner      string `bson:"owner"  json:"owner"`
}


func (app *App) GetImageName() string {
	imageName:=fmt.Sprintf("%v.%v:latest",app.Owner,app.Name)
	return imageName
}

func (app *App) GetPKGName() string {
	pkg:=app.GetImageName()
	pkgAddr:=fmt.Sprintf("%v/%v/%v/%v",constants.DockerRegistry,app.Owner,app.Name,pkg)
	return pkgAddr
}


