package handlers

import (
	"code-runner/internal/constants"
	"code-runner/internal/views"
	"github.com/dghubble/sessions"
)

var (
	sessionStore    = sessions.NewCookieStore([]byte(constants.SessionSecret), nil)
	contentsDir     = "internal/views/contents"
	jsDir     		= "internal/views/js"
	userAccessViews = map[string]*views.View{
		"index": views.NewView(
			"base", contentsDir+"/index.gohtml",jsDir+"/noJs.gohtml"),
		"workspace": views.NewView(
			"base", contentsDir+"/workspace.gohtml",jsDir+"/workspaceJs.gohtml"),
		"token": views.NewView(
			"base", contentsDir+"/token.gohtml",jsDir+"/noJs.gohtml"),
		"updateWs": views.NewView(
			"emptyBase", contentsDir+"/updateWs.gohtml"),
	}
	appsViews = map[string]*views.View{
		"createApp": views.NewView(
			"base", contentsDir+"/createApp.gohtml",jsDir+"/noJs.gohtml"),
		"viewApp": views.NewView(
			"base", contentsDir+"/viewApp.gohtml",jsDir+"/viewJs.gohtml"),
		"viewAppLog": views.NewView(
			"emptyBase", contentsDir+"/viewAppLogs.gohtml",jsDir+"/noJs.gohtml"),
	}
)
