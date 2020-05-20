package main

import (
	"code-runner/internal/config"
	"code-runner/internal/handlers"
	"code-runner/internal/store"
	"context"
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
	store.ClientStore = store.InitMongoStore(context.Background())

	apiAddr:=fmt.Sprintf("%v:%v",config.GetConfig().ApiAddress,config.GetConfig().ApiPort)
	server := &http.Server{
		Handler:      handlers.NewApi(),
		Addr:         apiAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,

	}

	var err error
	if config.GetConfig().EnableTls {
		log.Printf("Starting TLS Server listening on https://%s\n", apiAddr)
		err = server.ListenAndServeTLS(config.GetConfig().TlsCertFile,config.GetConfig().TlsKeyFile)
	} else {
		log.Printf("Starting Server listening on http://%s\n", apiAddr)
		err = server.ListenAndServe()
	}

	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
