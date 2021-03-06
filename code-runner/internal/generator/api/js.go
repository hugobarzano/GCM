package api

import (
	"code-runner/internal/generator/api/commons"
	"code-runner/internal/generator/api/js"
	"encoding/json"
	"fmt"
	"log"
)

type javascript struct {
	customNature
}

func (g *javascript) Init() Generator {
	g.files = make(map[string][]byte)
	g.model = make(map[string]interface{})
	g.spec = make([]byte, 0)
	return g
}

func (g *javascript) GetFiles() map[string][]byte {
	return g.files
}

func (g *javascript) WithName(name string) Generator {
	if name == "" {
		g.name = fmt.Sprintf("app%v", "test")
	} else {
		g.name = name
	}
	return g
}

func (g *javascript) WithPort(port int) Generator {
	g.port = port
	return g
}

func (g *javascript) WithInputSpec(spec string) Generator {
	g.spec = []byte(spec)
	err := json.Unmarshal([]byte(g.spec), &g.model)
	if err != nil {
		log.Println(err.Error())
	}
	return g
}

func (g *javascript) GenerateApi(des string) {

	data := make(map[string]interface{})
	data["ui"] = `<a href="ui/" class="navbar-brand"> Checkout API UI</a>`
	data["name"] = g.name
	data["des"] = des
	g.files["spec/spec.yml"] = commons.GenerateApiSpecFile(
		g.name,
		"",
		g.model)
	g.files["templates/index.html"] = commons.GenerateIndex(data)
	g.files["api/index.js"] = js.GenerateApi()
	g.files["server.js"] = js.GenerateServer(g.port)
	g.files["package.json"] = js.GenerateDependencies(g.name)
}
