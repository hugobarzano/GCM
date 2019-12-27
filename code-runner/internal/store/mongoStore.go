package store

import (
	"code-runner/internal/config"
	"code-runner/internal/constants"
	"code-runner/internal/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

type MongoStore struct {
	Client *mongo.Client
}

func InitMongoStore(ctx context.Context) *MongoStore {
	clientOptions := options.Client().ApplyURI(config.GetConfig().MongoUri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &MongoStore{client}
}

func (dao *MongoStore) TestConnection(ctx context.Context) {

	err := dao.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the MongoDB", err)
	} else {
		log.Println("Connected to MongoDB!")
	}
}

func (dao *MongoStore) CreateWorkspace(ctx context.Context, ws *models.Workspace) (*models.Workspace, error) {

	collection := dao.Client.Database(constants.Database).Collection(constants.WorkspacesCollection)
	_, err := collection.InsertOne(ctx, ws)
	if err != nil {
		return nil, err
	}
	return ws, nil
}

func (dao *MongoStore) CreateApp(ctx context.Context, app *models.App) (*models.App, error) {

	app.Status = models.INIT
	collection := dao.Client.Database(constants.Database).Collection(constants.AppsCollection)
	_, err := collection.InsertOne(ctx, app)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func (dao *MongoStore) GetWorkspace(ctx context.Context, owner string) (*models.Workspace, error) {

	var workspace *models.Workspace
	collection := dao.Client.Database(constants.Database).Collection(constants.WorkspacesCollection)
	filter := bson.M{"_id": owner}
	documentReturned := collection.FindOne(ctx, filter)
	if err := documentReturned.Decode(&workspace); err != nil {
		return nil, err
	}
	apps, err := dao.GetApps(ctx, owner)
	if err != nil {
		return nil, err
	}
	workspace.Apps = apps
	return workspace, nil
}

func (dao *MongoStore) GetApp(ctx context.Context, owner, name string) (*models.App, error) {
	var app *models.App
	collection := dao.Client.Database(constants.Database).Collection(constants.AppsCollection)
	filter := bson.M{"_id": name, "owner": owner}
	documentReturned := collection.FindOne(ctx, filter)
	if err := documentReturned.Decode(&app); err != nil {
		return nil, err
	}
	return app, nil
}

func (dao *MongoStore) DeleteApp(ctx context.Context, owner, name string) error {
	collection := dao.Client.Database(constants.Database).Collection(constants.AppsCollection)
	filter := bson.M{"_id": name, "owner": owner}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}

func (dao *MongoStore) GetApps(ctx context.Context, owner string) ([]models.App, error) {
	var apps []models.App
	var app models.App
	collection := dao.Client.Database(constants.Database).Collection(constants.AppsCollection)
	filter := bson.M{"owner": owner}
	cursor, err := collection.Find(ctx, filter, nil)
	if err != nil {
		fmt.Println("Error finding")
		return nil, err
	}
	for cursor.Next(ctx) {
		err := cursor.Decode(&app)
		if err != nil {
			fmt.Println("Error decoding from cursor")
			return nil, err
		}
		apps = append(apps, app)
	}
	if err := cursor.Err(); err != nil {
		fmt.Println("Error cursor")
		return nil, err
	}
	defer cursor.Close(ctx)
	return apps, nil
}

func (dao *MongoStore) UpdateApp(ctx context.Context, app *models.App) (*models.App, error) {
	collection := dao.Client.Database(constants.Database).Collection(constants.AppsCollection)
	var updated *models.App
	query := bson.M{"_id": app.Name, "owner": app.Owner}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	change := bson.M{"$set": app}
	res := collection.FindOneAndUpdate(ctx, query, change, &opt)
	err := res.Decode(&updated)
	if err != nil {
		fmt.Println("Error decoding updated app")
		return nil, err
	}
	return updated, nil
}

//func (dao *MongoStore)UpdateApp(ctx context.Context,app *models.App) (*models.Workspace, error) {
//	collection := dao.Client.Database(constants.Database).Collection(constants.WorkspacesCollection)
//	ws,err:=dao.GetWorkspace(ctx,app.Owner)
//	if err != nil {
//		return nil, err
//	}
//	a:=ws.GetApp(app.Name)
//	if a!=nil{}
//	query := bson.M{"_id": app.Owner}
//	change := bson.M{"$push": bson.M{"apps": app}}
//	_, err := collection.UpdateOne(ctx, query, change)
//	if err != nil {
//		return nil, err
//	}
//	ws,err:=dao.GetWorkspace(ctx,app.Owner)
//	if err != nil {
//		return nil, err
//	}
//	return ws, nil
//}
