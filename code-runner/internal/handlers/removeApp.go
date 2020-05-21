package handlers

import (
	"code-runner/internal/constants"
	"code-runner/internal/deploy"
	"code-runner/internal/generator"
	"code-runner/internal/store"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func removeApp(w http.ResponseWriter, r *http.Request) {

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

		go func() {
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
				http.Error(w,
					fmt.Sprintf("error Initialize docker engine:%s", err.Error()),
					http.StatusInternalServerError)
			}

			err = dockerApp.ContainerStop(ctx)
			if err != nil {
				fmt.Printf("error stoping app container:%s", err.Error())
			}

			err = dockerApp.ContainerRemove(ctx)
			if err != nil {
				fmt.Printf("error removing app container:%s", err.Error())
			}

			genApp := generator.GenApp{
				App: appObj,
			}

			genApp.InitGit(ctx,accessToken)
			_, err = genApp.DeleteRepo(ctx)
			if err != nil {
				http.Error(w,
					fmt.Sprintf("error removing code repository:%s", err.Error()),
					http.StatusInternalServerError)
			}

			err = store.ClientStore.DeleteApp(ctx,user,app)
			if err != nil {
				fmt.Println("error removing app: " + err.Error())
			}
		}()
		http.Redirect(w, r, "/workspace", http.StatusFound)
	}

	http.NotFound(w, r)
	return
}
