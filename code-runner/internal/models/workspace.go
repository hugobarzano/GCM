package models

import (
	"code-runner/internal/constants"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Name       string `bson:"_id"  json:"name"`
	Repository string `bson:"repo" json:"repo"`
	Spec       string `bson:"spec" json:"spec"`
	Des   string `bson:"des" json:"des,omitempty"`
	Url        string `bson:"url"  json:"url"`
	Owner        string `bson:"owner"  json:"owner"`
}

type Workspace struct {
	Owner string `bson:"_id" json:"owner,required"`
	Des   string `bson:"des" json:"des,omitempty"`
	Apps  []App  `bson:"apps" json:"apps,omitempty"`
}

func toMongoDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
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
	_, err := collection.InsertOne(context.Background(), ws)

	if err != nil {
		return  nil, err
	}

	return ws, nil
}

func InsertAppWithinWorkspace(client *mongo.Client, ws *Workspace, app *App) (*Workspace, error) {
	collection := client.Database(constants.Database).Collection(constants.WorkspacesCollection)
	query := bson.M{"_id": ws.Owner}
	change := bson.M{"$push":bson.M{"apps":app}}
	_,err:=collection.UpdateOne(context.Background(),query,change)

	if err != nil {
		return  nil, err
	}

	ws.Apps=append(ws.Apps,*app)
	return ws,nil
}

func RemoveAppWithinWorkspace(client *mongo.Client, ws *Workspace, appName string) (*Workspace, error) {
	collection := client.Database(constants.Database).Collection(constants.WorkspacesCollection)
	query := bson.M{"_id": ws.Owner}
	change := bson.M{"$pull":bson.M{"apps":bson.M{"_id":appName}}}
	_,err:=collection.UpdateOne(context.Background(),query,change)

	if err != nil {
		return  nil, err
	}
	for i:=range ws.Apps{
		fmt.Println(ws.Apps[i])
		if ws.Apps[i].Name == appName{
			ws.Apps[i]=App{}
		}
	}
	return ws,nil
}
