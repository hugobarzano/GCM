package handlers

import (
	"code-runner/internal/constants"
	"code-runner/internal/deploy"
	"context"
	"log"

	//"code-runner/internal/deploy"
	"code-runner/internal/models"
	"code-runner/internal/store"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func stopAppHandler(w http.ResponseWriter, r *http.Request) {

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

		stopApp(user,app)

		http.Redirect(w, r, "/workspace", http.StatusFound)
	}
	return
}

func stopApp(user,app string)  {
	ctx:= context.Background()
	appObj, err := store.ClientStore.GetApp(ctx, user, app)
	if err != nil {
		log.Println(fmt.Sprintf("error getting app:%s", err.Error()))
	}

	dockerApp := deploy.DockerApp{
		App: appObj,
	}

	err = dockerApp.Initialize()
	if err != nil {
		log.Println(fmt.Sprintf("error Initialize docker engine:%s", err.Error()))
	}

	err = dockerApp.ContainerStop(ctx)
	if err != nil {
		log.Println(fmt.Sprintf("error stoping app container: %s", err.Error()))
	}

	err = dockerApp.ContainerRemove(ctx)
	if err != nil {
		log.Println(fmt.Sprintf("error removing app container: %s", err.Error()))
	}

	appObj.Status = models.STOPPED
	appObj.Url = ""
	appObj.Spec["dockerId"] = ""

	_, err = store.ClientStore.UpdateApp(ctx, appObj)
	if err != nil {
		log.Println(fmt.Sprintf("error updating DB with stopped app:%s", err.Error()))
	}

}