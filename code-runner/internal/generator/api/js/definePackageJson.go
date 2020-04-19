package js

const JsPackageJson = `
{
  "name": "{{.}}",
  "version": "1.0.0",
  "description": "Generated {{.}} app",
  "main": "server.js",
  "scripts": {
    "test": "make test",
    "start": "node server.js"
  },
  "keywords": [
    "crud",
    "api",
	"dynamic api"
  ],
  "author": "bot-generator",
  "license": "Apache-2.0",
  "dependencies": {
    "express": "^4.17.1",
    "fs": "0.0.1-security",
    "swagger-routes-express": "^3.1.2",
    "swagger-ui-express": "^4.1.4",
    "yamljs": "^0.3.0"
  }
}
`
