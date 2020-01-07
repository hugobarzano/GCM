package main

import (
	"code-runner/internal/config"
	"code-runner/internal/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {


	port,ok:=os.LookupEnv("PORT")
	if ok{
		config.GetConfig().ApiPort=port
	}
	apiAddr:=fmt.Sprintf("%v:%v",config.GetConfig().ApiAddress,config.GetConfig().ApiPort)
	server := &http.Server{
		Handler:      handlers.NewApi(),
		Addr:         apiAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Starting Server listening on %s\n", apiAddr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
