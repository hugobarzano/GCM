package commons

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"
)

const (
	ApiBool    = "boolean"
	ApiInt     = "integer"
	ApiFloat   = "number"
	ApiDefault = "string"
)

type ApiSpec struct {
	Title     string
	ApiName   string
	ModelName string
	ModelData string
	ApiPkg    string
}

func GenerateApiSpecFile(name, pkg string, data map[string]interface{}) []byte {
	tpl, err := template.New("Spec").Parse(SpecTemplate)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	api := ApiSpec{
		Title:     getTitle(name),
		ApiName:   getApiName(name),
		ModelName: getModelName(name),
		ApiPkg:    pkg,
		ModelData: string(generateModelProperties(data)),
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, api)
	if err != nil {
		log.Fatalln(err)
	}
	return buf.Bytes()
}

func getTitle(name string) string {
	return strings.Title(fmt.Sprintf("API-%v generated", name))
}

func getModelName(name string) string {
	return strings.ToUpper(name)
}

func getApiName(name string) string {
	return strings.ToLower(name)
}

func generateModelProperties(spec map[string]interface{}) []byte {
	tpl, err := template.New("ModelProperties").Parse(ModelPropertiesTemplate)
	var buf bytes.Buffer
	err = tpl.Execute(&buf, string(generateProperties(spec)))
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return buf.Bytes()
}

func generateProperties(spec map[string]interface{}) []byte {
	tpl, err := template.New("Properties").Parse(PropertyTemplate)
	data := make(map[string]string)
	for key, value := range spec {
		data[key] = typeToSpec(value)
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return buf.Bytes()
}

func typeToSpec(input interface{}) string {

	switch input.(type) {
	case bool:
		return ApiBool
	case int:
		return ApiInt
	case float64, float32:
		return ApiFloat
	default:
		return ApiDefault
	}
}
