package handlers

import (
	"code-runner/internal/config"
	"code-runner/internal/constants"
	"code-runner/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"code-runner/internal/mongo"
	"code-runner/internal/views"
	"fmt"
	"github.com/dghubble/gologin/v2/github"
	oauth2Login "github.com/dghubble/gologin/v2/oauth2"
	"github.com/dghubble/sessions"
	"net/http"
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

func index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	if isAuthenticated(req) {
		http.Redirect(w, req, "/workspace", http.StatusFound)
		return
	}

	if err := userAccessViews["index"].Render(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func workspace(w http.ResponseWriter, req *http.Request) {

	session,err := sessionStore.Get(req,constants.SessionName)
	if err !=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	owner:=session.Values[constants.SessionUserName].(string)

	workspace,err:=models.GetWorkspace(
		databaseClient,
		bson.M{"owner": owner})

	//if err !=nil{ //Preguntar a ivan como controlar error
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	if workspace == nil{
		fmt.Println("First login for: "+owner)
		workspace,err=models.CreateWorkspace(databaseClient,&models.Workspace{
			Owner:owner,
		})
		if err !=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}else {
		fmt.Println(owner+ "Already has a workspace")
	}

	if err := userAccessViews["workspace"].Render(w, workspace); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func logout(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		sessionStore.Destroy(w, constants.SessionName)
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

func requireLogin(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, req *http.Request) {
		if !isAuthenticated(req) {
			http.Redirect(w, req, "/", http.StatusFound)
			return
		}
		next.ServeHTTP(w, req)
	}
	return http.HandlerFunc(fn)
}

func setupSession() http.Handler {

	fn := func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		githubUser, err := github.UserFromContext(ctx)
		fmt.Print(githubUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		githubToken, err := oauth2Login.TokenFromContext(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//Repository creation with token
		//ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken.AccessToken})
		//tc := oauth2.NewClient(ctx, ts)
		//client := googleGithub.NewClient(tc)
		//
		//r := &googleGithub.Repository{Name: googleGithub.String("generated-repo-33"),
		//	Private: googleGithub.Bool(false), Description: googleGithub.String("des")}
		//repo, _, err := client.Repositories.Create(ctx, "", r)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//fmt.Printf("Successfully created new repo: %v\n", repo.GetName())

		session := sessionStore.New(constants.SessionName)
		session.Values[constants.SessionUserKey] = *githubUser.ID
		session.Values[constants.SessionUserName] = *githubUser.Login
		session.Values[constants.SessionUserToken] = githubToken.AccessToken


		if err = session.Save(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, req, "/workspace", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}

func isAuthenticated(req *http.Request) bool {
	if _, err := sessionStore.Get(req, constants.SessionName); err == nil {
		return true
	}
	return false
}
