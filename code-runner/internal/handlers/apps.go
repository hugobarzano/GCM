package handlers

import (
	"code-runner/internal/constants"
	"code-runner/internal/models"
	"code-runner/internal/repos"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
)

func createApp(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()

	if err := req.ParseForm(); err != nil {
		http.Error(w,
			fmt.Sprintf("Error parsing form %s", err.Error()),
			http.StatusInternalServerError)
		return
	}

	session, err := sessionStore.Get(req, constants.SessionName)

	if err != nil {
		http.Error(w,
			fmt.Sprintf("Session Error: %s", err.Error()),
			http.StatusInternalServerError)
		return
	}

	app := models.App{
		Name: strings.Replace(
			req.FormValue("name"), "\"", "", -1),
		Des: strings.Replace(
			req.FormValue("description"), "\"", "", -1),
	}

	accessToken := session.Values[constants.SessionUserToken].(string)
	githubClient := repos.NewGithubClient(ctx, accessToken)
	appRepo := githubClient.CreateRepo(
		ctx,
		app.Name,
		app.Des)
	app.Repository=appRepo.GetURL()
	user := session.Values[constants.SessionUserName].(string)
	workspace, err := models.GetWorkspace(databaseClient, bson.M{"_id": user})
	_, err = models.InsertAppWithinWorkspace(databaseClient, workspace, app)

	if err != nil {
		http.Error(w,
			fmt.Sprintf("InsertAppWithinWorkspace Error: %s", err.Error()),
			http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "/workspace", http.StatusFound)
}

func removeApp(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		ctx := r.Context()
		vars := mux.Vars(r)
		app := vars["app"]
		app = strings.Replace(app, "\"", "", -1)
		fmt.Println(app)
		if app == "" {
			http.Error(w,
				fmt.Sprintf("Url Param 'app' is missing"),
				http.StatusInternalServerError)
			return
		}

		session, _ := sessionStore.Get(r, constants.SessionName)
		user := session.Values[constants.SessionUserName].(string)
		accessToken := session.Values[constants.SessionUserToken].(string)
		githubClient := repos.NewGithubClient(ctx, accessToken)
		res, err := githubClient.DeleteRepo(ctx, user, app)

		if err != nil {
			http.Error(w,
				fmt.Sprintf("error removing code repository:%s", err.Error()),
				http.StatusInternalServerError)
		}
		fmt.Println(res)

		workspace, _ := models.GetWorkspace(databaseClient, bson.M{"_id": user})
		_, err = models.RemoveAppWithinWorkspace(databaseClient, workspace, app)
		if err != nil {
			fmt.Println("error removing app: " + err.Error())
		}
		http.Redirect(w, r, "/workspace", http.StatusFound)
	}
	http.NotFound(w, r)
	return
}

//func test(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("method:", r.Method) //get request method
//	if r.Method == "GET" {
//		t, _ := template.ParseFiles("internal/views/contents/workspace.gohtml")
//		t.Execute(w, nil)
//	} else {
//		r.ParseForm()
//		// logic part of log in
//		fmt.Println("username:", r.Form["username"])
//		fmt.Println("password:", r.Form["password"])
//	}
//}
