package models

import (
	"code-runner/internal/constants"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	ID         string `bson:"_id" json:"_id,omitempty"`
	Repository string `json:"repository"`
	Url        string `json:"url"`
}


type Workspace struct {
	ID    string `bson:"_id" json:"_id,omitempty"`
	Owner string `bson:"owner" json:"owner,required"`
	Apps  []App  `bson:"apps" json:"apps,omitempty"`
}



func GetWorkspace(client *mongo.Client, filter bson.M) (*Workspace,error) {
	var workspace *Workspace
	collection := client.Database(constants.Database).Collection(constants.WorkspacesCollection)
	documentReturned := collection.FindOne(context.TODO(), filter)
	if err:=documentReturned.Decode(&workspace);err!=nil{
		return nil, err
	}
	return workspace, nil
}

func CreateWorkspace(client *mongo.Client, ws *Workspace) (*Workspace,error) {
	collection := client.Database(constants.Database).Collection(constants.WorkspacesCollection)
	insertResult, err := collection.InsertOne(context.Background(), ws)
	if err != nil {
		return  nil, err
	}
	ws.ID=insertResult.InsertedID.(string)
	return ws, nil
}

