package main

import (
	"code-runner/internal/constants"
	"code-runner/internal/handlers"
	"code-runner/internal/mongo"
	"log"
	"net/http"
	"time"
)

func init()  {
	mongo.TestMongoConnection()
}

func main() {

	server := &http.Server{
		Handler:      handlers.NewApi(),
		Addr:         constants.HttpAddress,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Starting Server listening on %s\n", constants.HttpAddress)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
