package repos

import (
	"context"
	"fmt"
	googleGithub "github.com/google/go-github/github"
	"golang.org/x/oauth2"
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

func (client *Client)CommitFile(ctx context.Context,path,repo string,
	opts *googleGithub.RepositoryContentFileOptions)  {

	_, _, err := client.github.Repositories.CreateFile(ctx,
		"", repo,
		path, opts)
	if err != nil {
		fmt.Println("Error on commit:"+err.Error())
	}
}

func (client *Client)DeleteRepo(ctx context.Context,owner,repo string)(
	*googleGithub.Response, error){

	res, err := client.github.Repositories.Delete(ctx,owner,repo)

	if err != nil {
		fmt.Println("Error Deleting repo"+err.Error())
		return nil,err
	}
	return res,nil
}
