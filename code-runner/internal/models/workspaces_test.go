package models

import (
	"code-runner/internal/config"
	"code-runner/internal/mongo"
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
	"testing"
	"time"
)
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateString(length int) string {
	return StringWithCharset(length, charset)
}

func Test_CreateAppWithinWorkspace(t *testing.T) {
	tesOwner:=GenerateString(10)
	databaseClient := mongo.GetClient(config.GetConfig().MongoUri)
	generator.New().Request.GetProtoFile()
	workspaceCreate,err:=CreateWorkspace(databaseClient,
		&Workspace{
			Owner: tesOwner,
			Apps:  []App{},
			Des: "testing",
		})

	if err !=nil {
		t.Errorf(err.Error())
		t.Fail()
	}
	fmt.Println(workspaceCreate)

	workspaceGet,err:=GetWorkspace( databaseClient, bson.M{"_id": tesOwner})
	fmt.Println(err)
	fmt.Println(workspaceGet)

	newApp:=&App{
		Name:"appName",
		Repository:"http://github.user.repo",
		Url: "TBD",
		Spec:"TBD",
		Owner: tesOwner,
	}

	workspaceWithApp,err := PushApp(databaseClient,workspaceGet,newApp)

	fmt.Println(workspaceWithApp)
	fmt.Println(err)

	newApp.Name="appNameToDele"
	workspaceWithApp,err = PushApp(databaseClient,workspaceGet,newApp)

	workspaceWithoutApp,err:= PopApp(databaseClient,workspaceWithApp,newApp.Name)

	fmt.Println(workspaceWithoutApp)
	fmt.Println(err)
	}
