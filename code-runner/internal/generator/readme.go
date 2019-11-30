package generator

import (
	"code-runner/internal/constants"
	"code-runner/internal/models"
)


func GenerateAppReadme(app *models.App)[]byte{
	readmeDoc:="# "+app.Name+"\n\n" + "## Description\n"+app.Des+"\n"
	readmeDoc=readmeDoc+constants.GeneratedBanner
	return []byte(readmeDoc)
}
