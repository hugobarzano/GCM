package devops

import (
	"fmt"
	"io/ioutil"
)

func GenerateInitGroovy()[]byte{
	fileData, err := ioutil.ReadFile("internal/resources/services/jenkins/init.groovy")
	if err != nil {
		fmt.Println("Error Reading")
		fmt.Println(err)
	}
	return fileData
}

func GeneratePluginsFile()[]byte{
	ciFileData, err := ioutil.ReadFile("internal/resources/services/jenkins/plugins.txt")
	if err != nil {
		fmt.Println("Error Reading")
		fmt.Println(err)
	}
	return ciFileData
}
