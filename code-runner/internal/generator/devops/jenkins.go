package devops

import (
	"io/ioutil"
	"log"
)

func GenerateInitGroovy() []byte {
	fileData, err := ioutil.ReadFile("internal/resources/services/jenkins/init.groovy")
	if err != nil {
		log.Println("Error Reading")
		log.Println(err)
	}
	return fileData
}

func GeneratePluginsFile() []byte {
	ciFileData, err := ioutil.ReadFile("internal/resources/services/jenkins/plugins.txt")
	if err != nil {
		log.Println("Error Reading")
		log.Println(err)
	}
	return ciFileData
}
