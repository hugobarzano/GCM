package main

import (
	"code-runner/internal/config"
	"code-runner/internal/handlers"
	"code-runner/internal/store"
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


	store.SinglePageImg = store.LoadImg("internal/store/resources/onePage.png")
	store.ApiRestImg = store.LoadImg("internal/store/resources/apiRest.png")
	store.DataServiceImg = store.LoadImg("internal/store/resources/dataService.png")
	store.DevOpsServiceImg = store.LoadImg("internal/store/resources/devOps.png")

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
