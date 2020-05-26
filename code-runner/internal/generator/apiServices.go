package generator

import (
	"code-runner/internal/generator/api"
	"strconv"
)

func (app *GenApp) generateApiService() {
	tech := app.App.Spec["tech"]
	var genNature api.Nature
	switch tech {
	case "go":
		genNature = api.Go
	case "python":
		genNature = api.Python
	case "js":
		genNature = api.JS
	}

	port, _ := strconv.Atoi(app.App.Spec["port"])
	apiGenerator := api.New(genNature)
	apiGenerator.Init().
		WithName(app.App.Name).
		WithPort(port).
		WithInputSpec(app.App.Spec["modelJson"]).GenerateApi(app.App.Des)
	app.Data = apiGenerator.GetFiles()
}
