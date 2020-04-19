package api

import (
	"code-runner/internal/generator/api/commons"
	"code-runner/internal/generator/api/py"
	"encoding/json"
	"fmt"
	"log"
)

type python struct {
	customNature
}

func (g *python) Init() Generator {
	g.files = make(map[string][]byte)
	g.model = make(map[string]interface{})
	g.spec = make([]byte, 0)
	return g
}

func (g *python) GetFiles() map[string][]byte			  {
	return g.files
}


func (g *python) WithName(name string) Generator {
	if name == "" {
		g.name = fmt.Sprintf("app%v", "test")
	} else {
		g.name = name
	}
	return g
}

func (g *python) WithPort(port int) Generator {
	g.port = port
	return g
}

func (g *python) WithInputSpec(spec string) Generator {
	g.spec = []byte(spec)
	err := json.Unmarshal(g.spec, &g.model)
	if err != nil {
		log.Println(err.Error())
	}
	return g
}

func (g *python) GenerateApi(des string) {

	data := make(map[string]interface{})
	data["ui"] = `<a href="api/ui" class="navbar-brand"> Checkout API UI</a>`
	data["name"]=g.name
	data["des"]=des
	g.files["spec/spec.yml"] = commons.GenerateApiSpecFile(
		g.name,
		"api.",
		g.model)
	g.files["templates/index.html"] = commons.GenerateIndex(data)
	g.files["api.py"] = py.GenerateApi()
	g.files["server.py"] = py.GenerateServer(g.port)
	g.files["requirements.txt"] = py.GenerateDependencies()

}
