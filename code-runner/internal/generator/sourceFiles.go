package generator

import 	googleGithub "github.com/google/go-github/github"

func BuilFileOptions(commit,user string, fileContent []byte ) *googleGithub.RepositoryContentFileOptions  {

	opts:=&googleGithub.RepositoryContentFileOptions{
		Message:   googleGithub.String(commit),
		Content:   fileContent,
		Branch:    googleGithub.String("master"),
		Committer: &googleGithub.CommitAuthor{
			Name: googleGithub.String(user),
			Email: googleGithub.String("mail.com")},
	}
	return opts
}
