package handlers

import (
	"code-runner/internal/config"
	"github.com/dghubble/gologin/v2"
	"github.com/dghubble/gologin/v2/github"
	"golang.org/x/oauth2"
	githubOAuth2 "golang.org/x/oauth2/github"
	"net/http"
)

func NewApi() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.Handle("/workspace", requireLogin(http.HandlerFunc(workspace)))
	mux.HandleFunc("/logout", logout)

	oauth2Config := &oauth2.Config{
		ClientID:     config.GetConfig().GithubClientID,
		ClientSecret: config.GetConfig().GithubClientSecret,
		RedirectURL:  "http://localhost:8080/github/callback",
		Endpoint:     githubOAuth2.Endpoint,
		Scopes:       []string{"public_repo"},
	}

	stateConfig := gologin.DebugOnlyCookieConfig
	mux.Handle("/github/login", github.StateHandler(
		stateConfig, github.LoginHandler(oauth2Config,nil)))
	mux.Handle("/github/callback", github.StateHandler(
		stateConfig, github.CallbackHandler(oauth2Config, setupSession(),nil)))
	return mux
}
