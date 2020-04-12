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
	img             = views.LoadImg("internal/views/resources/img2.png")
	userAccessViews = map[string]*views.View{
		"index": views.NewView(
			"base", contentsDir+"/index.gohtml",jsDir+"/noJs.gohtml",
			contentsDir + "/navbar/indexNavbar.gohtml"),
		"workspace": views.NewView(
			"base",
			contentsDir +"/workspace.gohtml",
			jsDir+"/workspaceJs.gohtml",
			contentsDir + "/navbar/workspaceNavbar.gohtml"),
		"token": views.NewView(
			"base", contentsDir+"/token.gohtml",jsDir+"/tokenJs.gohtml",
			contentsDir + "/navbar/tokenNavbar.gohtml"),
		"getWs": views.NewView(
			"emptyBase", contentsDir+"/workspace.gohtml"),
	}
	appsViews = map[string]*views.View{
		"createApp": views.NewView(
			"base", contentsDir+"/createApp.gohtml",jsDir+"/noJs.gohtml",contentsDir + "/navbar/noNavbar.gohtml"),
		"viewApp": views.NewView(
			"base", contentsDir+"/viewApp.gohtml",jsDir+"/viewJs.gohtml",contentsDir + "/navbar/viewAppNavbar.gohtml"),
		"getApp": views.NewView(
			"emptyBase", contentsDir+"/getApp.gohtml",contentsDir + "/navbar/noNavbar.gohtml"),
	}

 	upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	}
)
