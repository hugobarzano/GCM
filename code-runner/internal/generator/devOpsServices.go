package generator

import (
	"code-runner/internal/generator/devops"
)

func (app *GenApp) generateJenkinsService() {
	app.Data = make(map[string][]byte)
	app.Data["config/init.groovy"] = devops.GenerateInitGroovy()
	app.Data["config/plugins.txt"] = devops.GeneratePluginsFile()
}
