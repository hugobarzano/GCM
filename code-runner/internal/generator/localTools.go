package generator

import (
	"code-runner/internal/generator/local"
	"context"
	googleGithub "github.com/google/go-github/github")


func (app *GenApp) generateLocalTools() {
	app.Data = make(map[string][]byte)
	app.Data["makefile"] = local.GenMakefile()
	app.Data["bin/_buildImage.sh"] = local.GenbuildImage()
	app.Data["bin/_run.sh"] = local.GenRun(app.App.Spec["port"])
	app.Data["bin/_test.sh"] = local.GenTest()
	app.Data["bin/_pushImage.sh"] = local.GenPushImage()
	app.Data["bin/_pullImage.sh"] = local.GenPullImage()
}


func (app *GenApp) pushLocalTools(ctx context.Context,user,mail string)error {

	var commitMsg string
	var fileOptions *googleGithub.RepositoryContentFileOptions
	var err error

	for file,content := range app.Data{
		commitMsg="Generating "+file
		fileOptions = BuildFileOptions(commitMsg, user, mail,content)
		_, err = app.CommitFile(ctx, file, fileOptions)
		if err !=nil{
			return err
		}
	}
	return nil
}