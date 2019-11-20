package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"log"
	"testing"

	"golang.org/x/oauth2"
)

func Test(t *testing.T) {
	// In this example we're creating a new file in a repository using the
	// Contents API. Only 1 file per commit can be managed through that API.

	// Note that authentication is needed here as you are performing a modification
	// so you will need to modify the example to provide an oauth client to
	// github.NewClient() instead of nil. See the following documentation for more
	// information on how to authenticate with the client:
	// https://godoc.org/github.com/google/go-github/github#hdr-Authentication
	client := github.NewClient(nil)

	ctx := context.Background()
	fileContent := []byte("This is the content of my file\nand the 2nd line of it")

	// Note: the file needs to be absent from the repository as you are not
	// specifying a SHA reference here.
	opts := &github.RepositoryContentFileOptions{
		Message:   github.String("This is my commit message"),
		Content:   fileContent,
		Branch:    github.String("master"),
		Committer: &github.CommitAuthor{Name: github.String("yo"), Email: github.String("hugobarzano@gmail.com")},
	}
	_, _, err := client.Repositories.CreateFile(ctx, "users/hugobarzano", "GCM", "myNewFile.md", opts)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Test_CreateRpo(t *testing.T)  {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "d11bd4147f16759fa1c86675cf4f2948d010c3f8"})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	r := &github.Repository{Name: github.String("Test"), Private: github.Bool(true), Description: github.String("des")}
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo: %v\n", repo.GetName())
}

func Test_exampleRepositoriesService_GetReadme(t *testing.T) {
	client := github.NewClient(nil)

	readme, _, err := client.Repositories.GetReadme(context.Background(), "hugobarzano", "GCM", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	content, err := readme.GetContent()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("google/go-github README:\n%v\n", content)
}

func Test_ExampleUsersServiceListAll(t *testing.T) {
	client := github.NewClient(nil)
	opts := &github.UserListOptions{}
	for {
		users, _, err := client.Users.ListAll(context.Background(), opts)
		if err != nil {
			log.Fatalf("error listing users: %v", err)
		}
		if len(users) == 0 {
			break
		}
		opts.Since = *users[len(users)-1].ID
		// Process users...
	}
}
