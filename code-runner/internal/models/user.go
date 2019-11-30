package models

type User struct {
	Name       *string `json:"name"`
	Mail       *string `json:"mail"`
	GithubID   *string `json:"githubId"`
	GithubToken *string `json:"githubToken"`
}
