package generator

import (
	"context"
	"fmt"
googleGithub "github.com/google/go-github/github")


func (app *GenApp)generateSourceCode()  {

	switch nature := app.App.Spec["nature"]; nature {
	case "staticApp":
		app.generateStaticAppCode()
	case "mongodb":
		fmt.Println("TBD config generation.")
	default:
		fmt.Printf("NOT SUPPORTED")
	}
}

func (app *GenApp)pushSourceCode(ctx context.Context,user string)error {

	var commitMsg string
	var fileOptions *googleGithub.RepositoryContentFileOptions
	var err error

	for file,content := range app.Data{
		commitMsg="Generating "+file
		fileOptions = BuilFileOptions(commitMsg, user, content)
		_, err = app.CommitFile(ctx, file, fileOptions)
		if err !=nil{
			return err
		}
	}
	return nil
}
