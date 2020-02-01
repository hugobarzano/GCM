package handlers

import (
	"code-runner/internal/constants"
	"code-runner/internal/views"
	"github.com/dghubble/sessions"
	"github.com/gorilla/websocket"
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
		"getWs": views.NewView(
			"emptyBase", contentsDir+"/getWs.gohtml"),
	}
	appsViews = map[string]*views.View{
		"createApp": views.NewView(
			"base", contentsDir+"/createApp.gohtml",jsDir+"/noJs.gohtml"),
		"viewApp": views.NewView(
			"base", contentsDir+"/viewApp.gohtml",jsDir+"/viewJs.gohtml"),
		"getApp": views.NewView(
			"emptyBase", contentsDir+"/getApp.gohtml"),
	}

 	upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	}
)
