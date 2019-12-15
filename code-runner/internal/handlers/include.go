package handlers

import (
	"code-runner/internal/constants"
	"code-runner/internal/views"
	"github.com/dghubble/sessions"
)

var (
	sessionStore    = sessions.NewCookieStore([]byte(constants.SessionSecret), nil)
	contentsDir     = "internal/views/contents"
	userAccessViews = map[string]*views.View{
		"index": views.NewView(
			"base",
			contentsDir+"/index.gohtml"),
		"workspace": views.NewView(
			"base",
			contentsDir+"/workspace.gohtml"),
		"generate": views.NewView(
			"base",
			contentsDir+"/generate.gohtml"),
		"token": views.NewView(
			"base",
			contentsDir+"/token.gohtml"),
	}
)
