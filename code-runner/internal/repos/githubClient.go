package repos

import (
	"code-runner/internal/models"
	"code-runner/internal/tools"
	"context"
	"fmt"
	googleGithub "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type GitApp struct {
	App *models.App
	Github *googleGithub.Client
}

func getTokenSource (accessToken string) oauth2.TokenSource  {
	return oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: accessToken})
}

func (app *GitApp) Init(ctx context.Context, accessToken string){
	source:=getTokenSource(accessToken)
	authClient := oauth2.NewClient(ctx, source)
	app.Github= googleGithub.NewClient(authClient)
}


func(app *GitApp) GetSHA(ctx context.Context)string {
	commmit,_,_:=app.Github.Repositories.GetCommit(
		ctx,
		app.App.Owner,app.App.Name,"master")
	return 	commmit.GetSHA()[0:7]
}

func (app *GitApp)CreateRepo(ctx context.Context)*googleGithub.Repository{
	repoSpec := &googleGithub.Repository{
		Name:        googleGithub.String(app.App.Name),
		Private:     googleGithub.Bool(false),
		Description: googleGithub.String(app.App.Des),
	}
	repo, _, err := app.Github.Repositories.Create(ctx, "", repoSpec)
	if err != nil {
		fmt.Println("Error creating repo"+err.Error())
		return nil
	}
	return repo
}

func (app *GitApp)CreateFile(ctx context.Context,path string,
	fileOpts *googleGithub.RepositoryContentFileOptions)  (*googleGithub.RepositoryContentResponse,error){
	repoResponse,_, err := app.Github.Repositories.CreateFile(ctx,
		app.App.Owner, app.App.Name, path, fileOpts)
	if err != nil {
		fmt.Println("Error creating file"+err.Error())
		return nil,err
	}
	return repoResponse,err
}

func (app *GitApp)CommitFile(ctx context.Context,path string,
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

func (app *GitApp)GetFile(ctx context.Context,path string)(*string,error) {

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


func (app *GitApp)DeleteRepo(ctx context.Context)(
	*googleGithub.Response, error){
	res, err := app.Github.Repositories.Delete(ctx,app.App.Owner,app.App.Name)
	if err != nil {
		fmt.Println("Error Deleting repo"+err.Error())
		return nil,err
	}
	return res,nil
}

func (app *GitApp)GetTar(ctx context.Context)error {

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
//func NewGithubClient(ctx context.Context, accessToken string) Client  {
//	source:=getTokenSource(accessToken)
//	tokenClient := oauth2.NewClient(ctx, source)
//	return Client{ github:googleGithub.NewClient(tokenClient)}
//}
//
//func getTokenSource (accessToken string) oauth2.TokenSource  {
//	return oauth2.StaticTokenSource(&oauth2.Token{
//		AccessToken: accessToken})
//}
//func(client *Client) GetSha(ctx context.Context,owner string,repo string)string {
//	commmit,_,_:=client.github.Repositories.GetCommit(ctx,owner,repo,"master")
//	fmt.Print("COMMIT:\n")
//	fmt.Print(commmit.GetSHA())
//	return 	commmit.GetSHA()[0:7]
//}
//
//func (client *Client)CreateRepo(ctx context.Context,name,des string)*googleGithub.Repository{
//	repoSpec := &googleGithub.Repository{
//		Name:        googleGithub.String(name),
//		Private:     googleGithub.Bool(false),
//		Description: googleGithub.String(des),
//		}
//	repo, _, err := client.github.Repositories.Create(ctx, "", repoSpec)
//	if err != nil {
//		fmt.Println("Error creating repo"+err.Error())
//		return nil
//	}
//
//	return repo
//}
//
//func (client *Client)CreateFile(ctx context.Context,user,repo,path string,
//	opts *googleGithub.RepositoryContentFileOptions)  (*googleGithub.RepositoryContentResponse,error){
//
//		responseCon,res, err := client.github.Repositories.CreateFile(ctx, user, repo, path, opts)
//	if err != nil {
//		fmt.Println(err)
//		return nil,err
//	}
//	fmt.Println(responseCon)
//	fmt.Println(res)
//	return responseCon,err
//}
//
//func (client *Client)CommitFile(ctx context.Context,user,repo,path string,
//	opts *googleGithub.RepositoryContentFileOptions)  (*googleGithub.RepositoryContentResponse,error){
//
//	// if file exist we need to send sha attribute to override file
//
//	content,dir,res,err:=client.github.Repositories.GetContents(ctx, user, repo, path, nil)
//	fmt.Println("GET CONTENT:")
//	fmt.Println(content)
//	fmt.Println(res)
//	fmt.Println(dir)
//	if err!=nil{
//		if res.StatusCode != http.StatusNotFound{
//			return nil,err
//		}
//	}
//
//	if sha:=content.GetSHA();sha!=""{
//		opts.SHA=googleGithub.String(sha)
//	}
//
//	fmt.Println("SHA:")
//	fmt.Println(opts.SHA)
//	fmt.Println("CreateFILE")
//	responseCon,res, err := client.github.Repositories.CreateFile(ctx, user, repo, path, opts)
//	if err != nil {
//		fmt.Println(err)
//		return nil,err
//	}
//	fmt.Println(responseCon)
//	fmt.Println(res)
//	return responseCon,err
//}
//
//func (client *Client)GetFile(ctx context.Context,user,repo,path string)string  {
//
//	// if file exist we need to send sha attribute to override file
//
//	content,_,_,err:=client.github.Repositories.GetContents(ctx, user, repo, path, nil)
//
//	if err != nil{
//		fmt.Printf("Error geting file from repo: %v",err)
//	}
//
//	file,err:=content.GetContent()
//
//	if err != nil{
//		fmt.Printf("Error geting file content from repo: %v",err)
//	}
//	return file
//}
//
//func (client *Client)GetRepoTar(ctx context.Context,app models.App)error {
//
//	url, _, err := client.github.Repositories.GetArchiveLink(
//		ctx,
//		app.Owner,
//		app.Name,
//		googleGithub.Tarball,
//		&googleGithub.RepositoryContentGetOptions{
//			Ref: "master",
//		})
//	fmt.Println(url.String())
//	if err != nil{
//		fmt.Printf("Error geting ArchiveUrl from repo: %v",err)
//		return err
//	}
//
//	appLocalPath:=tools.GetAppLocalPath(app)
//	err=tools.DownloadFile(url.String(),appLocalPath)
//	if err != nil{
//		fmt.Printf("Error Downloadding TAR from repo: %v",err)
//		return err
//	}
//	return nil
//}
//
//func (client *Client)DeleteRepo(ctx context.Context,owner,repo string)(
//	*googleGithub.Response, error){
//	res, err := client.github.Repositories.Delete(ctx,owner,repo)
//
//	if err != nil {
//		fmt.Println("Error Deleting repo"+err.Error())
//		return nil,err
//	}
//	return res,nil
//}
