package commons

import (
	"bytes"
	"log"
	"text/template"
)

func GenerateIndex(uiLink string) []byte {
	tpl, err := template.New("Index").Parse(IndexTemplate)
	var buf bytes.Buffer
	err = tpl.Execute(&buf, uiLink)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return buf.Bytes()
}
