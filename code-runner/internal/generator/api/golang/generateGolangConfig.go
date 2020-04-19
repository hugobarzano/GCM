package golang

import (
	"bytes"
	"log"
	"text/template"
)

func GenerateDependencies(data map[string]interface{}) []byte {
	tpl, err := template.New("Modules").Parse(GoModulesTemplate)
	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return buf.Bytes()
}
