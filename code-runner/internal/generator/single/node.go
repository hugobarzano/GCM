package single


func GenPackageJson(name,des,owner string) []byte {
	pk := `
{
  "name": "`+name+`",
  "version": "1.0.0",
  "description": "`+des+`",
  "author": "`+owner+`",
  "main": "server.js",
  "scripts": {
    "start": "node server.js"
  },
  "dependencies": {
    "express": "^4.17.1",
    "fs": "0.0.1-security"
  }
}
`
	return []byte(pk)
}

func GenServerJs(port string) []byte {
	index := ` 
'use strict';
const express = require('express');
const fs = require('fs');

// Constants
const PORT = `+port+`;
const HOST = '0.0.0.0';

// App
const app = express();
app.get('/', (req, res) => {
    res.writeHead(200, {
        'Content-Type': 'text/html'
    });
    fs.readFile('./templates/index.html', null, function (error, data) {
        if (error) {
            res.status(404);
            res.write('ERROR');
        } else {
            res.write(data);
        }
        res.end();
    });
});

app.listen(PORT, HOST);
console.log("Running on http://"+HOST+":"+PORT);
`
	return []byte(index)
}

