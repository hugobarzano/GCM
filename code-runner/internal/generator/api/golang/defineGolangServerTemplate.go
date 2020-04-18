package golang

const GoServerTemplate = `
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"text/template"
	"{{.api}}/api"
)

func home(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			fmt.Fprintf(w, "Unable to load template")
		}
		t.Execute(w, "")
}

func main() {

	api.Init()
	router := mux.NewRouter()
	basePath:="api"
	apiPath:="{{.api}}"
	route:=fmt.Sprintf("/%v/%v",basePath,apiPath)
	router.HandleFunc("/", home).Methods(http.MethodGet)
	router.HandleFunc(route, api.ListAll).Methods(http.MethodGet)
	router.HandleFunc(route, api.CreateOne).Methods(http.MethodPost)
	router.HandleFunc(route+"/{id}", api.GetOne).Methods(http.MethodGet)
	router.HandleFunc(route+"/{id}", api.UpdateOne).Methods(http.MethodPut)
	router.HandleFunc(route+"/{id}", api.DeleteOne).Methods(http.MethodDelete)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	fmt.Println("Listening on http://0.0.0.0:{{.port}} !")
	log.Fatal(http.ListenAndServe("0.0.0.0:{{.port}}", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}`
