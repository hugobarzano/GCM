package handlers

import (
	"code-runner/internal/constants"
	"code-runner/internal/generator"
	"code-runner/internal/models"
	"code-runner/internal/repos"
	"fmt"
	"github.com/gorilla/mux"
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
	accessToken := session.Values[constants.SessionUserToken].(string)
	user := session.Values[constants.SessionUserName].(string)

	app := &models.App{
		Name: strings.Replace(
			req.FormValue("name"), "\"", "", -1),
		Des: strings.Replace(
			req.FormValue("description"), "\"", "", -1),
		Owner: user,
	}

	githubClient := repos.NewGithubClient(ctx, accessToken)
	appRepo := githubClient.CreateRepo(
		ctx,
		app.Name,
		app.Des)

	app.Repository = appRepo.GetCloneURL()
	workspace, err := models.GetWorkspace(databaseClient, user)
	_, err = models.PushApp(databaseClient, workspace, app)

	if err != nil {
		http.Error(w,
			fmt.Sprintf("PushApp Error: %s", err.Error()),
			http.StatusInternalServerError)
		return
	}

	readme := generator.GenerateAppReadme(app)
	fileOptions := repos.BuilFileOptions("Starting app...", user, "com.mail", readme)
	_,err=githubClient.CommitFile(ctx, user, app.Name, "README.md", fileOptions)
	if err != nil {
		http.Error(w,
			fmt.Sprintf("CommitFile Error: %s", err.Error()),
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

		workspace, _ := models.GetWorkspace(databaseClient, user)
		_, err = models.PopApp(databaseClient, workspace, app)
		if err != nil {
			fmt.Println("error removing app: " + err.Error())
		}
		http.Redirect(w, r, "/workspace", http.StatusFound)
	}
	http.NotFound(w, r)
	return
}

func generateApp(w http.ResponseWriter, r *http.Request) {

	switch method := r.Method; method {
	case http.MethodGet:
		fmt.Println("GET")
		vars := mux.Vars(r)
		app := vars["app"]
		app = strings.Replace(app, "\"", "", -1)
		fmt.Println(app)
		session, _ := sessionStore.Get(r, constants.SessionName)
		user := session.Values[constants.SessionUserName].(string)

		appObj,err:=models.GetApp(databaseClient,user,app)
		if err!=nil{
			http.Error(w,
				fmt.Sprintf("Error getting app %s", err.Error()),
				http.StatusNotFound)
			return
		}
		if err := userAccessViews["generate"].Render(w, appObj); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err := r.ParseForm(); err != nil {
			http.Error(w,
				fmt.Sprintf("Error parsing form %s", err.Error()),
				http.StatusInternalServerError)
			return
		}

		fmt.Println(r.FormValue("nature"))

		vars := mux.Vars(r)
		app := vars["app"]
		app = strings.Replace(app, "\"", "", -1)
		fmt.Println(app)
		session, _ := sessionStore.Get(r, constants.SessionName)
		user := session.Values[constants.SessionUserName].(string)
		accessToken := session.Values[constants.SessionUserToken].(string)
		appObj,err:=models.GetApp(databaseClient,user,app)
		if err!=nil{
			fmt.Println("error getting app")
			http.Error(w,
				fmt.Sprintf("Error getting app %s", err.Error()),
				http.StatusNotFound)
			return
		}
		dockerfile := generator.GenerateApacheDockerfile(appObj)
		fileOptions := repos.BuilFileOptions("Generating Dockerfile", user, "com.mail", dockerfile)
		ctx:=r.Context()
		githubClient := repos.NewGithubClient(ctx, accessToken)
		_,err=githubClient.CommitFile(ctx, user, appObj.Name, "Dockerfile", fileOptions)
		if err!= nil{
			fmt.Println("error getting tar")
		}
		err=githubClient.GetRepoTar(ctx,*appObj)
		if err!= nil{
			fmt.Println("error getting tar")
		}


		sha:=githubClient.GetSha(ctx,appObj.Owner,appObj.Name)
		fmt.Print("REPO BODY:")
		fmt.Print(sha)

		err =deployClient.BuildImage(ctx,*appObj,sha)
		if err!=nil{
			fmt.Println("error building imagen")
		}
		http.Redirect(w, r, "/workspace", http.StatusFound)
	default:
		fmt.Println("not supported")
	}
}
