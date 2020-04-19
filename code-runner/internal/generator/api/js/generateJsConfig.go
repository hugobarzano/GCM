package js

import (
	"bytes"
	"log"
	"text/template"
)

func GenerateDependencies(name string) []byte {
	tpl, err := template.New("PackageJson").Parse(JsPackageJson)
	var buf bytes.Buffer
	err = tpl.Execute(&buf, name)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return buf.Bytes()
}
