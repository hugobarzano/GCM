package models


type Workspace1 struct {
	ID        int     `json:"id" binding:"required"`
	Members   int     `json:"members"`
	Info 	  string  `json:"info"`
}

type App struct {
	ID     string     	 `bson:"_id" json:"_id,omitempty"`
	Url    string		 `bson:"url" json:"url,omitempty"`
}

type Workspace2 struct {
	ID        string     `json:"_id" json:"_id,omitempty"`
	Apps      []App      `json:"apps"`

}