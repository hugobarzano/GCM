package handlers

import (
	"code-runner/internal/config"
	"code-runner/internal/constants"
	"code-runner/internal/mongo"
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
	}

	databaseClient = mongo.GetClient(config.GetConfig().MongoUri)
)
