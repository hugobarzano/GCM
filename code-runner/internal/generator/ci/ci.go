package ci

import (
	"fmt"
	"io/ioutil"
)

func ImageBuilder()[]byte{
	ciFileData, err := ioutil.ReadFile("internal/resources/ci/imageBuilder.yml")
	if err != nil {
		fmt.Println("Error Reading")
		fmt.Println(err)
	}
	return ciFileData
}

func FastImageBuilder()[]byte{
	ciFileData, err := ioutil.ReadFile("internal/resources/ci/fastImageBuilder.yml")
	if err != nil {
		fmt.Println("Error Reading")
		fmt.Println(err)
	}
	return ciFileData
}