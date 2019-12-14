package models

type User struct {
	Name       	*string `json:"name"`
	Mail       	*string `json:"mail"`
	ID   		*int64 `json:"githubId"`
	AccessToken *string `json:"githubToken"`
}
