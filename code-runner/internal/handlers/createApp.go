package handlers

import (
	"code-runner/internal/constants"
	"code-runner/internal/deploy"
	"code-runner/internal/generator"
	"code-runner/internal/store"
	"fmt"
	"net/http"
)

func createApp(w http.ResponseWriter, r *http.Request) {

	switch method := r.Method; method {
	case http.MethodGet:
		if err := appsViews["createApp"].Render(w, ""); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:

		ctx:=r.Context()
		appObj,err:=getAppFromRequest(r)

		if err!=nil{
			http.Error(w,
				fmt.Sprintf("Input Error: %s", err.Error()),
				http.StatusInternalServerError)
			return
		}

		session, _ := sessionStore.Get(r, constants.SessionName)
		user := session.Values[constants.SessionUserName].(string)
		accessToken := session.Values[constants.SessionUserToken].(string)
		appObj.Owner=user

		genApp := generator.GenApp{
			App:appObj,
		}

		genApp.InitGit(ctx,accessToken)
		repo,err:=genApp.CreateRepo(ctx)

		if err!=nil{
			http.Error(w,
				fmt.Sprintf("Git Error: %s", err.Error()),
				http.StatusInternalServerError)
			return
		}

		appObj.Repository = repo.GetCloneURL()
		dao:=store.InitMongoStore(ctx)
		_,err=dao.CreateApp(ctx,appObj)
		if err != nil {
			http.Error(w,
				fmt.Sprintf("DB Error: %s", err.Error()),
				http.StatusInternalServerError)
			return
		}

		go genApp.InitializeCode(user,accessToken)

		dockerApp:=deploy.DockerApp{
			App:appObj,
		}
		go dockerApp.Start(accessToken)

		http.Redirect(w, r, "/workspace", http.StatusFound)
	default:
		fmt.Println("not supported")
	}
}

