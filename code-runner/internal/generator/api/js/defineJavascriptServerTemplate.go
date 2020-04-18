package js

const JsServerTemplate = `
const express = require('express')
const YAML = require('yamljs')
const swaggerUi = require('swagger-ui-express')
const { connector } = require('swagger-routes-express')
const api = require('./api')
const bodyParser = require('body-parser');
const fs = require('fs');

const home = (req, res) => {
    res.writeHead(200, {
        'Content-Type': 'text/html'
    });
    fs.readFile('./templates/index.html', null, function (error, data) {
        if (error) {
            res.status(404);
            res.write('Whoops!!');
        } else {
            res.write(data);
        }
        res.end();
    });
};


const makeApp = () => {
    const apiDefinition = YAML.load('./spec/spec.yml');
    const connect = connector(api, apiDefinition);
    const app = express();
    app.use('/api/ui', swaggerUi.serve, swaggerUi.setup(apiDefinition));
    app.use(bodyParser.urlencoded({ extended: true }));
    app.use(bodyParser.json());
    app.use("/home",home);
    connect(app);
    return app
};

makeApp().listen({{.}}, function () {
    console.log('Example app listening on port http://0.0.0.0:{{.}}!');
});`
