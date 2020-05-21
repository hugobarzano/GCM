package handlers

import (
	"code-runner/internal/constants"
	"code-runner/internal/models"
	"code-runner/internal/store"
	"fmt"
	"net/http"
)

func workspace(w http.ResponseWriter, req *http.Request) {

	session, err := sessionStore.Get(req, constants.SessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := session.Values[constants.SessionUserName].(string)
	ctx := req.Context()
	workspace, err := store.ClientStore.GetWorkspace(ctx, user)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}
	if workspace == nil {
		fmt.Println("First login for: " + user)
		workspace, err = store.ClientStore.CreateWorkspace(ctx, &models.Workspace{
			Owner: user,
			Des:   "Base workspace to app generation.",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		fmt.Println(user + ": Already has a workspace")
	}

	if err := userAccessViews["workspace"].Render(w, workspace); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func getWorkspace(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := req.FormValue("owner")
	if user == "" {
		http.Error(w, "Missing workspace owner", http.StatusInternalServerError)
		return
	}
	ctx := req.Context()
	workspace, err := store.ClientStore.GetWorkspace(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := userAccessViews["getWs"].Render(w, workspace); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
