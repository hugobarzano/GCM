package views

import (
	"html/template"
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
	log.Println("Rendering...")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func layoutFiles() []string {
	log.Println("Loading...")
	files, err := filepath.Glob(layoutsDir + "/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	return files
}
