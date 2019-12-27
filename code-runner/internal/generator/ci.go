package generator

import (
	"context"
	"fmt"
	"io/ioutil"
)

func readCI()[]byte{
	ciFileData, err := ioutil.ReadFile("internal/resources/ci/imageBuilder.yml")
	if err != nil {
		fmt.Println("Error Reading")
		fmt.Println(err)
	}
	return ciFileData
}


func (app *GenApp)generateCI(){
	app.CI=readCI()
}

func (app *GenApp)pushCI(ctx context.Context,user string)error{
	commit:="Generating CI workflow action to build docker image on push event..."
	ciFileOptions := BuilFileOptions(commit,user,app.CI)
	_,err:=app.CommitFile(ctx,".github/workflows/ci.yml", ciFileOptions)
	return err
}