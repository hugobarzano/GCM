package handlers

import (
	"bufio"
	"code-runner/internal/constants"
	"code-runner/internal/deploy"
	"code-runner/internal/store"
	"fmt"
	"github.com/gorilla/mux"
	"log"
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
		}
		return
	}
}


func viewAppLogSocket(w http.ResponseWriter, r *http.Request) {
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
	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer con.Close()

	appDocker:=deploy.DockerApp{
		App:nil,
	}
	err=appDocker.Initialize()
	if err != nil {
		http.Error(w,
			fmt.Sprintf("error Initialize docker engine:%s", err.Error()),
			http.StatusInternalServerError)
		return
	}
	out:=appDocker.GetContainerLogById2(ctx,dockerId)
	for {
		scanner := bufio.NewScanner(out)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
			err = con.WriteMessage(1, []byte(scanner.Text()))
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
		break
	}
}


func getApp(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		appName := r.FormValue("app")
		if appName == "" {
			http.Error(w, "Missing app Name field", http.StatusInternalServerError)
			return
		}

		session, _ := sessionStore.Get(r, constants.SessionName)
		user := session.Values[constants.SessionUserName].(string)

		dao := store.InitMongoStore(ctx)
		appObj, err := dao.GetApp(ctx, user, appName)
		if err != nil {
			http.Error(w,
				fmt.Sprintf("error getting app:%s", err.Error()),
				http.StatusInternalServerError)
		}

		if err := appsViews["getApp"].Render(w, appObj); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return

}

