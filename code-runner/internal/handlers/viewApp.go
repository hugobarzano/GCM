package handlers

import (
	"code-runner/internal/deploy"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Log struct {
	Data string
}
func viewAppLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ctx := r.Context()
		vars := mux.Vars(r)
		dockerId := vars["dockerId"]
		if dockerId == "" {
			http.Error(w,
				fmt.Sprintf("Url Param 'dockerId' is missing"),
				http.StatusInternalServerError)
			return
		}
		dockerApp:=deploy.DockerApp{
			App:nil,
		}

		if err:=dockerApp.Initialize();err!=nil{
			http.Error(w,
				err.Error(),
				http.StatusInternalServerError)
			return
		}


		logObj:=Log{
			Data:dockerApp.GetContainerLogById(ctx,dockerId),
		}


		if err := appsViews["viewApp"].Render(w, logObj); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}
}
