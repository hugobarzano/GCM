package generator

import (
	"code-runner/internal/models"
	"log"
	"testing"
)

func TestGenerateDockerfile(t *testing.T) {
	app := &models.App{
		Name:       "testApp",
		Repository: "https://github.repo.com",
		Des:        "some description",
		Owner:      "testOwner",
		Url:        "TBE",
	}
	properties := []dockerfileEntry{
		{"FROM", "httpd:2.4"},
		{"MAINTAINER", "testOwner"},
		{"COPY", "./html/ /usr/local/apache2/htdocs/"},
	}
	dockerfile := generateDockerfile(app, properties)
	log.Println(dockerfile)
	log.Println(string(dockerfile))
}
