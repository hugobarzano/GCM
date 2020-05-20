package handlers

import (
	"code-runner/internal/constants"
	"code-runner/internal/deploy"

	"code-runner/internal/store"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func runApp(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		ctx := r.Context()
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

		appObj, err := store.ClientStore.GetApp(ctx, user, app)
		if err != nil {
			http.Error(w,
				fmt.Sprintf("error getting app:%s", err.Error()),
				http.StatusInternalServerError)
		}

		dockerApp := deploy.DockerApp{
			App: appObj,
		}

		err = dockerApp.Initialize()
		if err != nil {
			fmt.Println("Initialize error: " + err.Error())
		}

		err = dockerApp.ContainerStop(ctx)
		if err != nil {
			fmt.Println("ContainerStop error: " + err.Error())
		}
		err=dockerApp.ContainerRemove(ctx)
		if err != nil {
			fmt.Println("ContainerRemove error: " + err.Error())
		}
		dockerApp.ContainerStart(accessToken)

		http.Redirect(w, r, "/workspace", http.StatusFound)
	}
	return
}
