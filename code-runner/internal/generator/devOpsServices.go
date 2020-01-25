package generator

import (
	"fmt"
	"io/ioutil"
)

func (app *GenApp) generateJenkinsService() {
	app.Data = make(map[string][]byte)
	app.Data["config/init.groovy"] = generateInitGroovy()
	app.Data["config/plugins.txt"] = generatePluginsFile()
}


func generateInitGroovy()[]byte{
	fileData, err := ioutil.ReadFile("internal/resources/services/jenkins/init.groovy")
	if err != nil {
		fmt.Println("Error Reading")
		fmt.Println(err)
	}
	return fileData
}

func generatePluginsFile()[]byte{
	ciFileData, err := ioutil.ReadFile("internal/resources/services/jenkins/plugins.txt")
	if err != nil {
		fmt.Println("Error Reading")
		fmt.Println(err)
	}
	return ciFileData
}