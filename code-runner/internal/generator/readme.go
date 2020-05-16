package generator

import (
	"code-runner/internal/constants"
	"context"
)

// ![Test and Push Docker Imagen](https://github.com/hugobarzano/aaa34/workflows/Test%20and%20Push%20Docker%20Imagen/badge.svg)

func (app *GenApp) generateReadme(){
	readmeDoc:="# "+app.App.Name+"\n\n" +
	"\n \n ![CI workflow](https://github.com/"+app.App.Owner+"/"+app.App.Name+"/workflows/Continuous%20Integration%20Workflow/badge.svg) \n"+
	"## Description\n"+app.App.Des+"\n"
	readmeDoc=readmeDoc+constants.GeneratedBanner
	app.Readme=[]byte(readmeDoc)
}

func (app *GenApp) pushReadme(ctx context.Context,user,mail string)error{
	readmeFileOptions := BuildFileOptions("Starting app...", user, mail,app.Readme)
	_, err := app.CommitFile(ctx, "README.md", readmeFileOptions)
	return err
}