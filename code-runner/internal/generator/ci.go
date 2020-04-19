package generator

import (
	"code-runner/internal/generator/ci"
	"context"
	googleGithub "github.com/google/go-github/github"
)



func (app *GenApp)generateCI(){
	app.CI = make(map[string][]byte)
	app.CI[".github/workflows/fastci.yml"]=ci.FastImageBuilder()
	app.CI[".github/workflows/ci.yml"]=ci.ImageBuilder()
}

func (app *GenApp)pushCI(ctx context.Context,user string)error {
	var commitMsg string
	var fileOptions *googleGithub.RepositoryContentFileOptions
	var err error
	for file,content := range app.CI{
		commitMsg="Generating CI workflow action "+file
		fileOptions = BuildFileOptions(commitMsg, user, content)
		_, err = app.CommitFile(ctx, file, fileOptions)
		if err !=nil{
			return err
		}
	}
	return nil
}