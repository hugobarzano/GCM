package generator

func (app *GenApp) generateNodeCode() {
	app.Data = make(map[string][]byte)
	app.Data["package.json"] = app.genPackageJson()
	app.Data["server.js"] = app.genServerJs()
}


func (app *GenApp) genPackageJson() []byte {
	pk := `
{
  "name": "`+app.App.Name+`",
  "version": "1.0.0",
  "description": "`+app.App.Des+`",
  "author": "`+app.App.Owner+`",
  "main": "server.js",
  "scripts": {
    "start": "node server.js"
  },
  "dependencies": {
    "express": "^4.16.1"
  }
}
`
	return []byte(pk)
}

func (app *GenApp) genServerJs() []byte {
	index := ` 
'use strict';
const express = require('express');

// Constants
const PORT = `+app.App.Spec["port"]+`;
const HOST = '0.0.0.0';

// App
const app = express();
app.get('/', (req, res) => {
  res.send('`+app.App.Name+`');
});

app.listen(PORT, HOST);
console.log("Running on http://${HOST}:${PORT}");
`
	return []byte(index)
}
