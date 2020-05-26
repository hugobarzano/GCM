package generator

import (
	"context"
	googleGithub "github.com/google/go-github/v31/github"
	"golang.org/x/oauth2"
	"log"
	"strconv"
	"strings"
)

func getTokenSource(accessToken string) oauth2.TokenSource {
	return oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: accessToken})
}

func (app *GenApp) InitGit(ctx context.Context, accessToken string) {
	source := getTokenSource(accessToken)
	authClient := oauth2.NewClient(ctx, source)
	app.Github = googleGithub.NewClient(authClient)
}

func (app *GenApp) GetLastRelease(ctx context.Context) string {

	tagList := []string{"0", "0", "0"}
	release, _, _ := app.Github.Repositories.GetLatestRelease(ctx, app.App.Owner, app.App.Name)
	if release != nil {
		tagList = strings.Split(*release.TagName, ".")
	}

	return strings.Join(tagList, ".")
}

func (app *GenApp) GetBeforeRelease(ctx context.Context) string {

	tagList := []string{"0", "0", "0"}
	release, _, err := app.Github.Repositories.GetLatestRelease(ctx, app.App.Owner, app.App.Name)
	if release != nil {
		tagList = strings.Split(*release.TagName, ".")
	}

	tagIndex := len(tagList) - 1
	tagInt, err := strconv.Atoi(tagList[tagIndex])
	if err != nil {
		log.Println(err.Error())
	}

	tagInt = tagInt - 1
	tagList[tagIndex] = strconv.Itoa(tagInt)
	return strings.Join(tagList, ".")
}

func (app *GenApp) GetNextRelease(ctx context.Context) string {

	tagList := []string{"0", "0", "0"}
	release, _, err := app.Github.Repositories.GetLatestRelease(ctx, app.App.Owner, app.App.Name)
	if release != nil {
		tagList = strings.Split(*release.TagName, ".")
	}

	tagIndex := len(tagList) - 1
	tagInt, err := strconv.Atoi(tagList[tagIndex])
	if err != nil {
		log.Println(err.Error())
	}

	tagInt = tagInt + 1
	tagList[tagIndex] = strconv.Itoa(tagInt)
	return strings.Join(tagList, ".")
}

//func (app *GenApp) CreateNextRelease(ctx context.Context)error{
//	release:=app.getNextRelease(ctx)
//	opt:=&googleGithub.RepositoryRelease{
//		TagName    : googleGithub.String(release),
//		TargetCommitish : googleGithub.String(""),
//		Name            :googleGithub.String(release),
//		Body           :googleGithub.String("Re-Generate source code release"),
//		Draft          :googleGithub.Bool(false),
//		Prerelease      :googleGithub.Bool(false),
//	}
//	rel,res,err:=app.Github.Repositories.CreateRelease(ctx,app.App.Owner,app.App.Name, opt)
//	log.Println(rel)
//	log.Println(res)
//	return err
//}

//func (app *GenApp) CreateInitialRelease(ctx context.Context)error{
//	opt:=&googleGithub.RepositoryRelease{
//		TagName    : googleGithub.String("0.0.0"),
//		TargetCommitish : googleGithub.String(""),
//		Name            :googleGithub.String("0.0.0"),
//		Body           :googleGithub.String("Initial source code release"),
//		Draft          :googleGithub.Bool(false),
//		Prerelease      :googleGithub.Bool(false),
//	}
//	rel,res,err:=app.Github.Repositories.CreateRelease(ctx,app.App.Owner,app.App.Name, opt)
//	log.Println(rel)
//	log.Println(res)
//	return err
//}

func (app *GenApp) GetSHA(ctx context.Context) string {
	commmit, _, _ := app.Github.Repositories.GetCommit(
		ctx,
		app.App.Owner, app.App.Name, "master")
	return commmit.GetSHA()[0:7]
}

func (app *GenApp) CreateRepo(ctx context.Context) (*googleGithub.Repository, error) {
	repoSpec := &googleGithub.Repository{
		Name:        googleGithub.String(app.App.Name),
		Private:     googleGithub.Bool(false),
		Description: googleGithub.String(app.App.Des),
	}
	repo, _, err := app.Github.Repositories.Create(ctx, "", repoSpec)
	if err != nil {
		log.Println("error creating repo: " + err.Error())
		return nil, err
	}
	repo.GetReleasesURL()
	return repo, nil
}

func (app *GenApp) CreateFile(ctx context.Context, path string,
	fileOpts *googleGithub.RepositoryContentFileOptions) (*googleGithub.RepositoryContentResponse, error) {
	repoResponse, _, err := app.Github.Repositories.CreateFile(ctx,
		app.App.Owner, app.App.Name, path, fileOpts)
	if err != nil {
		log.Println("error creating file: " + err.Error())
		return nil, err
	}
	return repoResponse, err
}

func (app *GenApp) CommitFile(ctx context.Context, path string,
	fileOpts *googleGithub.RepositoryContentFileOptions) (
	*googleGithub.RepositoryContentResponse, error) {
	content, _, _, err := app.Github.Repositories.GetContents(ctx,
		app.App.Owner, app.App.Name, path, nil)
	if err != nil {
		log.Println(err.Error() + "...Generating")
	}
	if sha := content.GetSHA(); sha != "" {
		fileOpts.SHA = googleGithub.String(sha)
	}

	responseCon, _, err := app.Github.Repositories.CreateFile(ctx,
		app.App.Owner, app.App.Name, path, fileOpts)
	if err != nil {
		log.Println("error creating file: " + err.Error())
		return nil, err
	}
	return responseCon, err
}

func (app *GenApp) GetFile(ctx context.Context, path string) (*string, error) {

	content, _, _, err := app.Github.Repositories.GetContents(
		ctx, app.App.Owner, app.App.Name, path, nil)

	if err != nil {
		log.Println("Error geting file from repo: %v", err)
		return nil, err
	}
	file, err := content.GetContent()
	if err != nil {
		log.Println("Error geting file content from repo: %v", err)
		return nil, err
	}
	return &file, nil
}

func (app *GenApp) DeleteRepo(ctx context.Context) (
	*googleGithub.Response, error) {
	res, err := app.Github.Repositories.Delete(ctx, app.App.Owner, app.App.Name)
	if err != nil {
		log.Println("Error Deleting repo" + err.Error())
		return nil, err
	}
	return res, nil
}
