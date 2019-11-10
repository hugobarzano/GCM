package models


type Workspace struct {
	ID        int     `json:"id" binding:"required"`
	Members   int     `json:"members"`
	Info 	  string  `json:"info"`
}
