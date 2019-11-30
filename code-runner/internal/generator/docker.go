package generator

import (
	"code-runner/internal/models"
)

type  dockerfileEntry struct {
	Action string
	Data string
}

func generateDockerfile(app *models.App,properties []dockerfileEntry)[]byte{

	dockerfile:="# Dockerfile from "+app.Repository+"\n"
	for iterator :=range properties{
		dockerfile=dockerfile+properties[iterator].Action+"    "+properties[iterator].Data+"  \n"
	}
	return []byte(dockerfile)
}

func GenerateApacheDockerfile(app *models.App) 	[]byte {
	properties:=[]dockerfileEntry{
		{"FROM","httpd:2.4"},
		{"MAINTAINER", app.Owner},
		{"COPY","./html/ /usr/local/apache2/htdocs/"},
	}
	dockerfile:=generateDockerfile(app,properties)
	return []byte(dockerfile)
}
