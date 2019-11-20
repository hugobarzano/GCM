package models

import (
	"code-runner/internal/config"
	"code-runner/internal/mongo"
	"fmt"
	"testing"
)

func Test_CreateWorkspace(t *testing.T)   {
	databaseClient:= mongo.GetClient(config.GetConfig().MongoUri)
	workspace:=CreateWorkspace(databaseClient,
		&Workspace{
			ID: "testing3",
			Owner: "hugo",
		})

	//workspace:=GetWorkspace( databaseClient, bson.M{"owner": "hugo"})
	fmt.Println(workspace)

}
