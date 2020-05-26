package ci

import (
	"io/ioutil"
	"log"
)

func ImageBuilder() []byte {
	ciFileData, err := ioutil.ReadFile("internal/resources/ci/imageBuilder.yml")
	if err != nil {
		log.Println("Error Reading")
		log.Println(err)
	}
	return ciFileData
}

func FastImageBuilder() []byte {
	ciFileData, err := ioutil.ReadFile("internal/resources/ci/fastImageBuilder.yml")
	if err != nil {
		log.Println("Error Reading")
		log.Println(err)
	}
	return ciFileData
}
