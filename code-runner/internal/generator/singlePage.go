package generator

import "code-runner/internal/generator/single"

func (app *GenApp) generateApacheSinglePageCode() {
	app.Data = make(map[string][]byte)
	app.Data["html/index.html"] = single.GenIndexHtml(
		app.App.Name,
		app.App.Des)

	app.Data["html/js/index.js"] = single.GenIndexJs(app.App.Owner)
	app.Data["html/css/style.css"] = single.GenStyleCss()
}

func (app *GenApp) generateNodeSinglePageCode() {
	app.Data = make(map[string][]byte)
	app.Data["package.json"] = single.GenPackageJson(
		app.App.Name,
		app.App.Des,
		app.App.Owner)

	app.Data["templates/index.html"] = single.GenIndexHtml(
		app.App.Name,
		app.App.Des)

	app.Data["server.js"] = single.GenServerJs(app.App.Spec["port"])
}



