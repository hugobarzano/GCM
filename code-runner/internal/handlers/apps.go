package handlers

import (
	"code-runner/internal/constants"
	"code-runner/internal/models"
	"code-runner/internal/repos"
	"fmt"
	googleGithub "github.com/google/go-github/github"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func createApp(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()

	if err := req.ParseForm(); err != nil {
		http.Error(w,
			fmt.Sprintf("Error parsing form %s",err.Error()),
			http.StatusInternalServerError)
		return
	}

	session, err := sessionStore.Get(req, constants.SessionName)

	if err != nil {
		http.Error(w,
			fmt.Sprintf("Session Error: %s",err.Error()),
			http.StatusInternalServerError)
		return
	}

	accessToken := session.Values[constants.SessionUserToken].(string)
	githubClient := repos.NewGithubClient(ctx,accessToken)
	appRepo:=githubClient.CreateRepo(
		ctx,
		req.FormValue("name"),
		req.FormValue("description"),)
	app := models.App{
		Name:       googleGithub.Stringify(appRepo.Name),
		Repository: appRepo.GetURL(),
		Url:        "TBD",
		Spec: "TBD",
	}
	user:=session.Values[constants.SessionUserName].(string)
	workspace,err:=models.GetWorkspace( databaseClient, bson.M{"_id": user})
	_,err = models.InsertAppWithinWorkspace(databaseClient,workspace,app)

	if err != nil {
		http.Error(w,
			fmt.Sprintf("InsertAppWithinWorkspace Error: %s",err.Error()),
			http.StatusInternalServerError)
		return
	}

	http.Redirect(w, req, "/workspace", http.StatusFound)
}

func removeApp(w http.ResponseWriter, r *http.Request)  {

	if r.Method == http.MethodPut {
		fmt.Println(r.Body)
		app, ok := r.URL.Query()["app"]

		if !ok || len(app[0]) < 1 {
			log.Println("Url Param 'app' is missing")
			return
		}
		session, _ := sessionStore.Get(r, constants.SessionName)
		user:=session.Values[constants.SessionUserName].(string)
		workspace,_:=models.GetWorkspace( databaseClient, bson.M{"_id": user})
		fmt.Println(app)
		_,err := models.RemoveAppWithinWorkspace(databaseClient,workspace,app[0])
		if err!=nil{
			fmt.Println("error removing app: "+err.Error())

		}

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
