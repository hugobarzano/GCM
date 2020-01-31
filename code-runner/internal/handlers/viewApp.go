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

type Log struct {
	Data string
}
func viewApp(w http.ResponseWriter, r *http.Request) {
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

		if err := appsViews["viewApp"].Render(w, appObj); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
}

func viewAppLogs(w http.ResponseWriter, r *http.Request) {

	ctx:=r.Context()
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dockerId := r.FormValue("dockerId")
	if dockerId == "" {
		http.Error(w, "Missing dockerId", http.StatusInternalServerError)
		return
	}

	appDocker:=deploy.DockerApp{
		App:nil,
	}
	err:=appDocker.Initialize()
	if err != nil {
		http.Error(w,
			fmt.Sprintf("error Initialize docker engine:%s", err.Error()),
			http.StatusInternalServerError)
	}

	log:= struct {
		Line string
	}{
		Line:appDocker.GetContainerLogById(ctx,dockerId),
	}

	if err := appsViews["viewAppLog"].Render(w, log); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

