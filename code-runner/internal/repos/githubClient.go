package repos

import (
	"code-runner/internal/models"
	"code-runner/internal/tools"
	"context"
	"fmt"
	googleGithub "github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"net/http"
)

type Client struct {
	github *googleGithub.Client
}

func NewGithubClient(ctx context.Context, accessToken string) Client  {
	source:=getTokenSource(accessToken)
	tokenClient := oauth2.NewClient(ctx, source)
	return Client{ github:googleGithub.NewClient(tokenClient)}
}

func getTokenSource (accessToken string) oauth2.TokenSource  {
	return oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: accessToken})
}
func(client *Client) GetSha(ctx context.Context,owner string,repo string)string {
	commmit,_,_:=client.github.Repositories.GetCommit(ctx,owner,repo,"master")
	fmt.Print("COMMIT:\n")
	fmt.Print(commmit.GetSHA())
	return 	commmit.GetSHA()[0:7]
}

func (client *Client)CreateRepo(ctx context.Context,name,des string)*googleGithub.Repository{
	repoSpec := &googleGithub.Repository{
		Name:        googleGithub.String(name),
		Private:     googleGithub.Bool(false),
		Description: googleGithub.String(des),
	}
	repo, _, err := client.github.Repositories.Create(ctx, "", repoSpec)
	if err != nil {
		fmt.Println("Error creating repo"+err.Error())
		return nil
	}

	return repo
}

func (client *Client)CommitFile(ctx context.Context,user,repo,path string,
	opts *googleGithub.RepositoryContentFileOptions)  (*googleGithub.RepositoryContentResponse,error){

	// if file exist we need to send sha attribute to override file

	content,_,res,err:=client.github.Repositories.GetContents(ctx, user, repo, path, nil)

	if err!=nil{
		if res.StatusCode != http.StatusNotFound{
			return nil,err
		}
	}

	if sha:=content.GetSHA();sha!=""{
		opts.SHA=googleGithub.String(sha)
	}
	response, _, err := client.github.Repositories.CreateFile(ctx, user, repo, path, opts)
	if err != nil {
		return nil,err
	}
	return response,err
}

func (client *Client)GetFile(ctx context.Context,user,repo,path string)string  {

	// if file exist we need to send sha attribute to override file

	content,_,_,err:=client.github.Repositories.GetContents(ctx, user, repo, path, nil)

	if err != nil{
		fmt.Printf("Error geting file from repo: %v",err)
	}

	file,err:=content.GetContent()

	if err != nil{
		fmt.Printf("Error geting file content from repo: %v",err)
	}
	return file
}

func (client *Client)GetRepoTar(ctx context.Context,app models.App)error {

	url, _, err := client.github.Repositories.GetArchiveLink(
		ctx,
		app.Owner,
		app.Name,
		googleGithub.Tarball,
		&googleGithub.RepositoryContentGetOptions{
			Ref: "master",
		})
	fmt.Println(url.String())
	if err != nil{
		fmt.Printf("Error geting ArchiveUrl from repo: %v",err)
		return err
	}

	appLocalPath:=tools.GetAppLocalPath(app)
	err=tools.DownloadFile(url.String(),appLocalPath)
	if err != nil{
		fmt.Printf("Error Downloadding TAR from repo: %v",err)
		return err
	}
	return nil
}

func (client *Client)DeleteRepo(ctx context.Context,owner,repo string)(
	*googleGithub.Response, error){
	client.github.Repositories.Paackages
	res, err := client.github.Repositories.Delete(ctx,owner,repo)

	if err != nil {
		fmt.Println("Error Deleting repo"+err.Error())
		return nil,err
	}
	return res,nil
}
