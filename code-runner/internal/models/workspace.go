package models

import (
	"code-runner/internal/constants"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)



type Workspace struct {
	Owner string `bson:"_id" json:"owner,required"`
	Des   string `bson:"des" json:"des,omitempty"`
	Apps  []App  `bson:"apps" json:"apps,omitempty"`
}

func CreateWorkspace(client *mongo.Client, ws *Workspace) (*Workspace, error) {
	collection := client.Database(constants.Database).Collection(constants.WorkspacesCollection)
	_, err := collection.InsertOne(context.Background(), ws)

	if err != nil {
		return nil, err
	}

	return ws, nil
}

func GetWorkspace(client *mongo.Client, user string) (*Workspace, error) {
	var workspace *Workspace
	collection := client.Database(constants.Database).Collection(constants.WorkspacesCollection)
	filter := bson.M{"_id": user}
	documentReturned := collection.FindOne(context.TODO(), filter)

	if err := documentReturned.Decode(&workspace); err != nil {
		return nil, err
	}

	return workspace, nil
}

func GetApp(client *mongo.Client, owner,name string) (*App, error) {
	ws, err := GetWorkspace(client, owner)
	if err != nil {
		return nil, err
	}
	for iterator := range ws.Apps {
		if ws.Apps[iterator].Name == name {
			return &ws.Apps[iterator],nil
			break
		}
	}
	return nil, errors.New("App not found")
}

func PushApp(client *mongo.Client, ws *Workspace, app *App) (*Workspace, error) {
	collection := client.Database(constants.Database).Collection(constants.WorkspacesCollection)
	query := bson.M{"_id": ws.Owner}
	change := bson.M{"$push": bson.M{"apps": app}}
	_, err := collection.UpdateOne(context.Background(), query, change)

	if err != nil {
		return nil, err
	}

	ws.Apps = append(ws.Apps, *app)
	return ws, nil
}

func PopApp(client *mongo.Client, ws *Workspace, appName string) (*Workspace, error) {
	collection := client.Database(constants.Database).Collection(constants.WorkspacesCollection)
	query := bson.M{"_id": ws.Owner}
	change := bson.M{"$pull": bson.M{"apps": bson.M{"_id": appName}}}
	_, err := collection.UpdateOne(context.Background(), query, change)

	if err != nil {
		return nil, err
	}
	for i := range ws.Apps {
		fmt.Println(ws.Apps[i])
		if ws.Apps[i].Name == appName {
			ws.Apps[i] = App{}
		}
	}
	return ws, nil
}
