package api

import (
	"code-runner/internal/generator/api/commons"
	"code-runner/internal/generator/api/golang"
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/gojson"
	"log"
	"strings"
)

type goApi struct {
	customNature
	nativeModel []byte
}

func (g *goApi) Init() Generator {
	g.files = make(map[string][]byte)
	g.model = make(map[string]interface{})
	g.nativeModel = make([]byte, 0)
	g.spec = make([]byte, 0)
	return g
}

func (g *goApi) GetFiles() map[string][]byte {
	return g.files
}

func (g *goApi) WithName(name string) Generator {
	if name == "" {
		g.name = fmt.Sprintf("app%v", "test")
	} else {
		g.name = name
	}
	return g
}

func (g *goApi) WithPort(port int) Generator {
	g.port = port
	return g
}

func (g *goApi) WithInputSpec(spec string) Generator {
	g.spec = []byte(spec)
	err := json.Unmarshal([]byte(g.spec), &g.model)
	if err != nil {
		log.Println(err.Error())
	}

	specReader:=strings.NewReader(spec)
	modelGenerated, err := gojson.Generate(
		specReader,
		gojson.ParseJson,
		strings.ToUpper(g.name),
		"model", []string{"json", "yml"},
		true, true)
	if err != nil {
		log.Println(err.Error())
	}

	g.nativeModel = customizeModel(modelGenerated)
	return g
}

func customizeModel(inputModel []byte) []byte {
	stringModel := string(inputModel)
	IDField := "        ID string      `json:\"id\" yml:\"id\"` \n}"
	customModel := strings.Replace(stringModel,"}",IDField,1)
	return []byte(customModel)
}


func (g *goApi) GenerateApi(des string) {

	data := make(map[string]interface{})
	data["api"] = strings.ToLower(g.name)
	data["model"] = strings.ToUpper(g.name)
	data["port"] = g.port
	data["name"]=g.name
	data["des"]=des
	g.files["spec/spec.yml"] = commons.GenerateApiSpecFile(
		g.name,
		"api.",
		g.model)
	g.files["templates/index.html"] = commons.GenerateIndex(data)
	g.files["api/api.go"] = golang.GenerateApi(data)
	g.files["model/model.go"] = g.nativeModel
	g.files["server.go"] = golang.GenerateServer(data)
	g.files["go.mod"] = golang.GenerateDependencies(data)
}
