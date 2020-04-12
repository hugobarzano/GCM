package views

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

var (
	layoutsDir string = "internal/views/layouts"
)

type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, files ...string) *View {
	files = append(layoutFiles(), files...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	fmt.Println("Rendering...")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func layoutFiles() []string {
	fmt.Println("Loading...")
	files, err := filepath.Glob(layoutsDir + "/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}

func LoadImg(path string)string{
	imageFile, err := ioutil.ReadFile(path)

	if err != nil {
		log.Println(err.Error())
		return ""
	}
	encodedString := base64.StdEncoding.EncodeToString(imageFile)
	return encodedString
}