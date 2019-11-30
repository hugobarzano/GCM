package main

import (
	"code-runner/internal/constants"
	"code-runner/internal/handlers"
	"code-runner/internal/mongo"
	"log"
	"net/http"
)

func init()  {
	mongo.TestMongoConnection()
}

func main() {
	log.Printf("Starting Server listening on %s\n", constants.HttpAddress)
	err := http.ListenAndServe(constants.HttpAddress, handlers.NewApi())
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
