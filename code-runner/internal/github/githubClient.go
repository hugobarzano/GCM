package github

import (
	"context"
	googleGithub "github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func GetTokenClient(ctx context.Context, accessToken string) *googleGithub.Client  {
	source:=getTokenSource(accessToken)
	tokenClient := oauth2.NewClient(ctx, source)
	return googleGithub.NewClient(tokenClient)
}

func getTokenSource (accessToken string) oauth2.TokenSource  {
	return oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: accessToken})
}

