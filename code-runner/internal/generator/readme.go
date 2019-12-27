package generator

import (
	"code-runner/internal/constants"
	"code-runner/internal/models"
	"context"
)


func GenerateAppReadme(app *models.App)[]byte{
	readmeDoc:="# "+app.Name+"\n\n" + "## Description\n"+app.Des+"\n"
	readmeDoc=readmeDoc+constants.GeneratedBanner
	return []byte(readmeDoc)
}


func (app *GenApp) initReadme(){
	readmeDoc:="# "+app.App.Name+"\n\n" + "## Description\n"+app.App.Des+"\n"
	readmeDoc=readmeDoc+constants.GeneratedBanner
	app.Readme=[]byte(readmeDoc)
}

func (app *GenApp) pushReadme(ctx context.Context,user string)error{
	readmeFileOptions := BuilFileOptions("Starting app...", user, app.Readme)
	_, err := app.CommitFile(ctx, "README.md", readmeFileOptions)
	return err
}