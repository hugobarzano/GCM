package generator

import (
	"code-runner/internal/models"
	"fmt"
	"testing"
)

func TestGenerateDockerfile(t *testing.T) {
	app:=&models.App{
		Name:"testApp",
		Repository:"https://github.repo.com",
		Des:"some description",
		Owner: "testOwner",
		Spec: "TBD",
		Url: "TBE",
	}
	properties:=[]dockerfileEntry{
		{"FROM","httpd:2.4"},
		{"MAINTAINER", "testOwner"},
		{"COPY","./html/ /usr/local/apache2/htdocs/"},
	}
	dockerfile:=generateDockerfile(app,properties)
	fmt.Println(dockerfile)
	fmt.Println(string(dockerfile))
}
