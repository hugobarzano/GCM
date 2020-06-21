package handlers

import (
	"code-runner/internal/constants"
	"code-runner/internal/generator"
	"code-runner/internal/store"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

func runAppHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		vars := mux.Vars(r)
		app := vars["app"]
		app = strings.Replace(app, "\"", "", -1)
		if app == "" {
			http.Error(w,
				fmt.Sprintf("Url Param 'app' is missing"),
				http.StatusInternalServerError)
			return
		}
		session, _ := sessionStore.Get(r, constants.SessionName)
		user := session.Values[constants.SessionUserName].(string)
		accessToken := session.Values[constants.SessionUserToken].(string)

		go runApp(user, accessToken, app)

		http.Redirect(w, r, "/workspace", http.StatusFound)
	}
	return
}

func runApp(user, token, app string) {
	ctx := context.Background()
	appObj, err := store.ClientStore.GetApp(ctx, user, app)
	if err != nil {
		log.Println("error Getting App: " + err.Error())
	}

	if appObj != nil {
		dockerApp := generator.DockerApp{
			App: appObj,
		}
		err = dockerApp.Initialize()
		if err != nil {
			log.Println("Initialize error: " + err.Error())
		}

		err = dockerApp.ContainerStop(ctx)
		if err != nil {
			log.Println("ContainerStop error: " + err.Error())
		}
		err = dockerApp.ContainerRemove(ctx)
		if err != nil {
			log.Println("ContainerRemove error: " + err.Error())
		}
		dockerApp.ContainerStart(token)
	}
}
