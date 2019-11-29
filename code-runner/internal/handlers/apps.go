package handlers

import (
	"code-runner/internal/constants"
	"fmt"
	googleGithub "github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"html/template"
	"log"
	"net/http"
)

func createApp(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	if err := req.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := sessionStore.Get(req, constants.SessionName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	accessToken := session.Values[constants.SessionUserToken].(string)

	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: accessToken})
	tokenClient := oauth2.NewClient(ctx, tokenSource)
	githubClient := googleGithub.NewClient(tokenClient)

	r := &googleGithub.Repository{
		Name:        googleGithub.String(req.FormValue("name")),
		Private:     googleGithub.Bool(false),
		Description: googleGithub.String(req.FormValue("description"))}
	repo, _, err := githubClient.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully created repo: %v\n%v", repo.GetName(), repo.GetURL())

	//app := models.App{
	//	Name:       req.FormValue("name"),
	//	Repository: repo.GetURL(),
	//	Url:        "TBD",
	//}

	http.Redirect(w, req, "/workspace", http.StatusFound)

}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("internal/views/contents/workspace.gohtml")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
