package golang

import (
	"bytes"
	"log"
	"text/template"
)

func GenerateServer(data map[string]interface{}) []byte {
	tpl, err := template.New("Server").Parse(GoServerTemplate)
	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return buf.Bytes()
}
