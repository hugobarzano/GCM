package generator

import googleGithub "github.com/google/go-github/v31/github"

func BuildFileOptions(commit, user, mail string, fileContent []byte) *googleGithub.RepositoryContentFileOptions {

	opts := &googleGithub.RepositoryContentFileOptions{
		Message: googleGithub.String(commit),
		Content: fileContent,
		Branch:  googleGithub.String("master"),
		Committer: &googleGithub.CommitAuthor{
			Name:  googleGithub.String(user),
			Email: googleGithub.String(mail)},
	}
	return opts
}
