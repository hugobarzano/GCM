package handlers

import (
	"code-runner/internal/constants"
	"github.com/dghubble/gologin/v2/github"
	oauth2Login "github.com/dghubble/gologin/v2/oauth2"
	"net/http"
)



func index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	if isAuthenticated(req) {
		http.Redirect(w, req, "/workspace", http.StatusFound)
		return
	}

	if err := userAccessViews["index"].Render(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getToken(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodGet {
		session,_:=sessionStore.Get(req, constants.SessionName)
		token:= struct {
			Key string
		}{
			Key:session.Values[constants.SessionUserToken].(string),
		}

		if err := userAccessViews["token"].Render(w, token); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
	//http.Redirect(w, req, "/", http.StatusFound)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		sessionStore.Destroy(w, constants.SessionName)
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

func requireLogin(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, req *http.Request) {
		if !isAuthenticated(req) {
			http.Redirect(w, req, "/", http.StatusFound)
			return
		}
		next.ServeHTTP(w, req)
	}
	return http.HandlerFunc(fn)
}

func setupSession() http.Handler {

	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		githubUser, err := github.UserFromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		githubToken, err := oauth2Login.TokenFromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session := sessionStore.New(constants.SessionName)
		session.Values[constants.SessionUserKey] = *githubUser.ID
		session.Values[constants.SessionUserName] = *githubUser.Login
		session.Values[constants.SessionUserToken] = githubToken.AccessToken

		if err = session.Save(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, req, "/workspace", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

func isAuthenticated(req *http.Request) bool {
	if _, err := sessionStore.Get(req, constants.SessionName); err == nil {
		return true
	}
	return false
}
