package models

//
//import (
//	"code-runner/internal/config"
//	"fmt"
//	"github.com/golang/protobuf/protoc-gen-go/generator"
//	"testing"
//)
//
//
//func Test_CreateAppWithinWorkspace(t *testing.T) {
//	tesOwner:=GenerateString(10)
//	databaseClient := mongo.GetClient(config.GetConfig().MongoUri)
//	generator.New().Request.GetProtoFile()
//	workspaceCreate,err:=CreateWorkspace(databaseClient,
//		&Workspace{
//			Owner: tesOwner,
//			Apps:  []App{},
//			Des: "testing",
//		})
//
//	if err !=nil {
//		t.Errorf(err.Error())
//		t.Fail()
//	}
//	log.Println(workspaceCreate)
//
//	workspaceGet,err:=GetWorkspace( databaseClient, tesOwner)
//	log.Println(err)
//	log.Println(workspaceGet)
//
//	newApp:=&App{
//		Name:"appName",
//		Repository:"http://github.user.repo",
//		Url: "TBD",
//		Spec:"TBD",
//		Owner: tesOwner,
//	}
//
//	workspaceWithApp,err := PushApp(databaseClient,workspaceGet,newApp)
//
//	log.Println(workspaceWithApp)
//	log.Println(err)
//
//	newApp.Name="appNameToDele"
//	workspaceWithApp,err = PushApp(databaseClient,workspaceGet,newApp)
//
//	workspaceWithoutApp,err:= PopApp(databaseClient,workspaceWithApp,newApp.Name)
//
//	log.Println(workspaceWithoutApp)
//	log.Println(err)
//	}
