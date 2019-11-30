package generator

import (
	"code-runner/internal/constants"
	"code-runner/internal/models"
)

type readme struct {
	Title string
	Des string


}

func GenerateAppReadme(app *models.App)[]byte{
	readmeDoc:="# "+app.Name+"\n\n" + "## Description\n"+app.Des+"\n"
	readmeDoc=readmeDoc+constants.GeneratorBanner
	return []byte(readmeDoc)
}
