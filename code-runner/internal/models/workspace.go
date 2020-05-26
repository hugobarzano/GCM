package models

type Workspace struct {
	Owner string `bson:"_id" json:"owner,required"`
	Des   string `bson:"des" json:"des,omitempty"`
	Apps  []App  `json:"apps,omitempty"`
}
