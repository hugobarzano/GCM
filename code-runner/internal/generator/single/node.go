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
    "express": "^4.16.1"
  }
}
`
	return []byte(pk)
}

func GenServerJs(name, port string) []byte {
	index := ` 
'use strict';
const express = require('express');

// Constants
const PORT = `+port+`;
const HOST = '0.0.0.0';

// App
const app = express();
app.get('/', (req, res) => {
  res.send('`+name+`');
});

app.listen(PORT, HOST);
console.log("Running on http://${HOST}:${PORT}");
`
	return []byte(index)
}

