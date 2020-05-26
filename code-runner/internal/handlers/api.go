package handlers

import (
	"code-runner/internal/config"
	"fmt"
	"github.com/dghubble/gologin/v2"
	"github.com/dghubble/gologin/v2/github"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	githubOAuth2 "golang.org/x/oauth2/github"
	"net/http"
)

func NewApi() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/", index)
	mux.Handle("/workspace", requireLogin(http.HandlerFunc(workspace)))
	mux.Handle("/getWs", http.HandlerFunc(getWorkspace))
	mux.HandleFunc("/token", getToken).Methods(http.MethodGet)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/createApp", createApp)
	mux.HandleFunc("/remove/{app}", removeAppHandler)
	mux.HandleFunc("/stop/{app}", stopAppHandler)
	mux.HandleFunc("/run/{app}", runAppHandler)
	mux.HandleFunc("/regenerate/{app}", reGenerateAppHandler)
	mux.HandleFunc("/view/{app}", viewApp)
	mux.Handle("/getApp", http.HandlerFunc(getApp))
	mux.HandleFunc("/logsSocket", viewAppLogSocket)
	oauth2Config := &oauth2.Config{
		ClientID:     config.GetConfig().GithubClientID,
		ClientSecret: config.GetConfig().GithubClientSecret,
		RedirectURL:  fmt.Sprintf("%v/%v", config.GetConfig().ApiDns, "github/callback"),
		Endpoint:     githubOAuth2.Endpoint,
		Scopes: []string{
			"repo",
			"public_repo",
			"delete_repo",
			"workflow",
			"read:packages",
			"write:packages",
			"delete:packages"},
	}

	stateConfig := gologin.DebugOnlyCookieConfig
	mux.Handle("/github/login", github.StateHandler(
		stateConfig, github.LoginHandler(oauth2Config, nil)))
	mux.Handle("/github/callback", github.StateHandler(
		stateConfig, github.CallbackHandler(oauth2Config, setupSession(), nil)))
	return mux
}
