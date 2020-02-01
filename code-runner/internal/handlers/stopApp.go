package handlers

import (
	"code-runner/internal/constants"
	"code-runner/internal/deploy"

	//"code-runner/internal/deploy"
	"code-runner/internal/models"
	"code-runner/internal/store"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func stopApp(w http.ResponseWriter, r *http.Request) {

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
		dao := store.InitMongoStore(ctx)
		appObj, err := dao.GetApp(ctx, user, app)
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
			http.Error(w,
				fmt.Sprintf("error Initialize docker engine:%s", err.Error()),
				http.StatusInternalServerError)
		}

		err = dockerApp.ContainerStop(ctx)
		if err != nil {
				fmt.Printf("error stoping app container:%s", err.Error())
		}

		go dockerApp.ContainerRemove(ctx)

		appObj.Status = models.STOPPED
		appObj.Url = ""
		appObj.Spec["dockerId"] = ""

		_, err = dao.UpdateApp(ctx, appObj)
		if err != nil {
			http.Error(w,
				fmt.Sprintf("error updating DB with stopped app:%s", err.Error()),
				http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/workspace", http.StatusFound)
	}
	return
}
