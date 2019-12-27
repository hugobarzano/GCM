package generator

import (
	"code-runner/internal/tools"
	"context"
	"fmt"
	"golang.org/x/oauth2"
	googleGithub "github.com/google/go-github/github"
)

func getTokenSource (accessToken string) oauth2.TokenSource  {
	return oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: accessToken})
}

func (app *GenApp) InitGit(ctx context.Context, accessToken string){
	source:=getTokenSource(accessToken)
	authClient := oauth2.NewClient(ctx, source)
	app.Github= googleGithub.NewClient(authClient)
}


func(app *GenApp) GetSHA(ctx context.Context)string {
	commmit,_,_:=app.Github.Repositories.GetCommit(
		ctx,
		app.App.Owner,app.App.Name,"master")
	return 	commmit.GetSHA()[0:7]
}

func (app *GenApp)CreateRepo(ctx context.Context)(*googleGithub.Repository,error){
	repoSpec := &googleGithub.Repository{
		Name:        googleGithub.String(app.App.Name),
		Private:     googleGithub.Bool(false),
		Description: googleGithub.String(app.App.Des),
	}
	repo, _, err := app.Github.Repositories.Create(ctx, "", repoSpec)
	if err != nil {
		fmt.Println("Error creating repo"+err.Error())
		return nil,err
	}
	return repo,nil
}

func (app *GenApp)CreateFile(ctx context.Context,path string,
	fileOpts *googleGithub.RepositoryContentFileOptions)  (*googleGithub.RepositoryContentResponse,error){
	repoResponse,_, err := app.Github.Repositories.CreateFile(ctx,
		app.App.Owner, app.App.Name, path, fileOpts)
	if err != nil {
		fmt.Println("Error creating file"+err.Error())
		return nil,err
	}
	return repoResponse,err
}

func (app *GenApp)CommitFile(ctx context.Context,path string,
	fileOpts *googleGithub.RepositoryContentFileOptions)(
	*googleGithub.RepositoryContentResponse,error){
	content,_,_,err:=app.Github.Repositories.GetContents(ctx,
		app.App.Owner, app.App.Name, path, nil)
	if err!=nil{
		fmt.Println("Error Getting content file"+err.Error())
		//return nil,err
	}
	if sha:=content.GetSHA();sha!=""{
		fileOpts.SHA=googleGithub.String(sha)
	}

	responseCon,_, err := app.Github.Repositories.CreateFile(ctx,
		app.App.Owner, app.App.Name, path, fileOpts)
	if err != nil {
		fmt.Println("Error creating file"+err.Error())
		return nil,err
	}
	return responseCon,err
}

func (app *GenApp)GetFile(ctx context.Context,path string)(*string,error) {

	content, _, _, err := app.Github.Repositories.GetContents(
		ctx, app.App.Owner, app.App.Name, path, nil)

	if err != nil {
		fmt.Printf("Error geting file from repo: %v", err)
		return nil,err
	}
	file, err := content.GetContent()
	if err != nil {
		fmt.Printf("Error geting file content from repo: %v", err)
		return nil,err
	}
	return &file,nil
}


func (app *GenApp)DeleteRepo(ctx context.Context)(
	*googleGithub.Response, error){
	res, err := app.Github.Repositories.Delete(ctx,app.App.Owner,app.App.Name)
	if err != nil {
		fmt.Println("Error Deleting repo"+err.Error())
		return nil,err
	}
	return res,nil
}

func (app *GenApp)GetTar(ctx context.Context)error {

	url, _, err := app.Github.Repositories.GetArchiveLink(
		ctx,
		app.App.Owner,
		app.App.Name,
		googleGithub.Tarball,
		&googleGithub.RepositoryContentGetOptions{
			Ref: "master",
		})
	if err != nil{
		fmt.Printf("Error geting ArchiveUrl from repo: %v",err)
		return err
	}

	err=tools.DownloadFile(url.String(),app.App.GetLocalPath())
	if err != nil{
		fmt.Printf("Error Downloadding TAR from repo: %v",err)
		return err
	}
	return nil
}
