package generator

import (
	"context"
	"fmt"
	"io/ioutil"
	googleGithub "github.com/google/go-github/github"
)

func imageBuilder()[]byte{
	ciFileData, err := ioutil.ReadFile("internal/resources/ci/imageBuilder.yml")
	if err != nil {
		fmt.Println("Error Reading")
		fmt.Println(err)
	}
	return ciFileData
}

func fastImageBuilder()[]byte{
	ciFileData, err := ioutil.ReadFile("internal/resources/ci/fastImageBuilder.yml")
	if err != nil {
		fmt.Println("Error Reading")
		fmt.Println(err)
	}
	return ciFileData
}


func (app *GenApp)generateCI(){
	app.CI = make(map[string][]byte)
	app.CI[".github/workflows/fastci.yml"]=fastImageBuilder()
	app.CI[".github/workflows/ci.yml"]=imageBuilder()
}

func (app *GenApp)pushCI(ctx context.Context,user string)error {

	var commitMsg string
	var fileOptions *googleGithub.RepositoryContentFileOptions
	var err error

	for file,content := range app.CI{
		commitMsg="Generating CI workflow action "+file
		fileOptions = BuilFileOptions(commitMsg, user, content)
		_, err = app.CommitFile(ctx, file, fileOptions)
		if err !=nil{
			return err
		}
	}
	return nil
}