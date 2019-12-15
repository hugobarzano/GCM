package main

import (
	"code-runner/internal/constants"
	"code-runner/internal/handlers"
	"log"
	"net/http"
	"time"
)

func main() {

	server := &http.Server{
		Handler:      handlers.NewApi(),
		Addr:         constants.HttpAddress,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Starting Server listening on %s\n", constants.HttpAddress)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
