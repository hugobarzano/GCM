package py

import (
	"bytes"
	"log"
	"text/template"
)

func GenerateServer(port int) []byte {
	tpl, err := template.New("Server").Parse(PythonServerTemplate)
	var buf bytes.Buffer
	err = tpl.Execute(&buf, port)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return buf.Bytes()

}
