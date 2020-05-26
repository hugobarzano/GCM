package handlers

import (
	"code-runner/internal/constants"
	"code-runner/internal/generator"
	"code-runner/internal/store"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func createApp(w http.ResponseWriter, r *http.Request) {

	switch method := r.Method; method {
	case http.MethodGet:
		if err := appsViews["createApp"].Render(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:

		ctx := r.Context()
		reqApp, err := getAppFromRequest(r)

		if err != nil {
			http.Error(w,
				fmt.Sprintf("Input Error: %s", err.Error()),
				http.StatusInternalServerError)
			return
		}

		if ok := reqApp.validateRequest(); !ok {
			_ = appsViews["createApp"].Render(w, reqApp)
			return
		}

		session, _ := sessionStore.Get(r, constants.SessionName)
		user := session.Values[constants.SessionUserName].(string)
		mail := session.Values[constants.SessionUserMail].(string)
		accessToken := session.Values[constants.SessionUserToken].(string)
		reqApp.App.Owner = user

		genApp := generator.GenApp{
			App: reqApp.App,
		}

		genApp.InitGit(ctx, accessToken)
		repo, err := genApp.CreateRepo(ctx)

		if err != nil {
			if strings.Contains(err.Error(), "already exists on this account") {
				reqApp.Errors["Name"] = "There is another repository with this application name in your github account"
				_ = appsViews["createApp"].Render(w, reqApp)
			} else {
				http.Error(w,
					fmt.Sprintf("Git Error: %s", err.Error()),
					http.StatusInternalServerError)
			}
			return
		}

		reqApp.App.Repository = repo.GetCloneURL()
		_, err = store.ClientStore.CreateApp(ctx, reqApp.App)
		if err != nil {
			http.Error(w,
				fmt.Sprintf("DB Error: %s", err.Error()),
				http.StatusInternalServerError)
			return
		}

		go genApp.InitializeCode(user, accessToken, mail)

		http.Redirect(w, r, "/workspace", http.StatusFound)
	default:
		log.Println("not supported")
	}
}
