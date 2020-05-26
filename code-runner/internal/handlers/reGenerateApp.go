package handlers

import (
	"code-runner/internal/constants"
	"code-runner/internal/generator"
	"code-runner/internal/models"
	"code-runner/internal/store"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

func reGenerateAppHandler(w http.ResponseWriter, r *http.Request) {

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

		go reGenerateApp(user, accessToken, app)

		http.Redirect(w, r, "/workspace", http.StatusFound)
	}
}

func reGenerateApp(user, token, app string) {
	ctx := context.Background()
	appObj, err := store.ClientStore.GetApp(ctx, user, app)
	if err != nil {
		log.Println(fmt.Sprintf("error getting app: %s", err.Error()))
	}

	dockerApp := generator.DockerApp{
		App: appObj,
	}

	err = dockerApp.Initialize()

	if err != nil {
		log.Println(fmt.Sprintf("error Initialize docker engine:%s", err.Error()))
	}

	appObj.Status = models.STOPPED
	appObj.Url = ""

	_, err = store.ClientStore.UpdateApp(ctx, appObj)
	if err != nil {
		log.Println(fmt.Sprintf("error updating DB with stopped app:%s", err.Error()))
	}

	err = dockerApp.ContainerStop(ctx)
	if err != nil {
		log.Println(fmt.Sprintf("error stoping app container:%s", err.Error()))
	}

	err = dockerApp.ContainerRemove(ctx)
	if err != nil {
		log.Println(fmt.Sprintf("error removing app container:%s", err.Error()))
	}

	err = dockerApp.ImageRemove(ctx, token)
	if err != nil {
		log.Println(fmt.Sprintf("error removing image: %s", err.Error()))
	}

	genApp := generator.GenApp{
		App: appObj,
	}
	genApp.InitGit(ctx, token)
	genApp.DisableCI()
	if err := genApp.PushCI(ctx, user, user); err != nil {
		log.Println(fmt.Sprintf("PushFile Error: %s", err.Error()))
	}
	appObj.Status = models.INIT
	_, err = store.ClientStore.UpdateApp(ctx, appObj)
	if err != nil {
		log.Println(fmt.Sprintf("error updating DB with stopped app:%s", err.Error()))
	}
	genApp.ReGenerateCode(user, token, user)

}
