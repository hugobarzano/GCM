package generator

import (
	"code-runner/internal/models"
	"context"
	"fmt"
)

type  dockerfileEntry struct {
	Action string
	Data string
}

func generateDockerfile(app *models.App,properties []dockerfileEntry)[]byte{

	dockerfile:="# Dockerfile2 from "+app.Repository+"\n"
	for iterator :=range properties{
		dockerfile=dockerfile+properties[iterator].Action+"    "+properties[iterator].Data+"  \n"
	}
	return []byte(dockerfile)
}

func GenerateApacheDockerfile(app *models.App) 	[]byte {
	properties:=[]dockerfileEntry{
		{"FROM","httpd:2.4"},
		{"MAINTAINER", app.Owner},
		{"RUN","sed 's/^Listen 80/Listen "+app.Spec["port"]+"/g' /usr/local/apache2/conf/httpd.conf > httpd.new"},
		{"RUN","mv httpd.new /usr/local/apache2/conf/httpd.conf"},
		{"COPY","html/ /usr/local/apache2/htdocs/"},
		{"EXPOSE", app.Spec["port"]},
	}
	dockerfile:=generateDockerfile(app,properties)
	return []byte(dockerfile)
}

func GenerateMongoDockerfile(app *models.App) 	[]byte {
	properties:=[]dockerfileEntry{
		{"FROM","mongo:3.6"},
		{"MAINTAINER", app.Owner},
		{"COPY","./config/mongod.conf /etc/mongod.conf"},
		{"EXPOSE", app.Spec["port"]},
		{"ENTRYPOINT", "[\"mongod\", \"-f\", \"/etc/mongod.conf\"]"},
	}
	dockerfile:=generateDockerfile(app,properties)
	return []byte(dockerfile)
}

func GenerateMysqlDockerfile(app *models.App) 	[]byte {
	properties:=[]dockerfileEntry{
		{"FROM","mysql:8.0"},
		{"MAINTAINER", app.Owner},
		{"ENV", "MYSQL_ALLOW_EMPTY_PASSWORD=1"},
		{"COPY","./config/mysql.cnf /etc/mysql/conf.d/mysql.cnf"},
		{"EXPOSE", app.Spec["port"]},
	}
	dockerfile:=generateDockerfile(app,properties)
	return []byte(dockerfile)
}

func GenerateRedisDockerfile(app *models.App) 	[]byte {
	properties:=[]dockerfileEntry{
		{"FROM","redis:6.0-rc"},
		{"MAINTAINER", app.Owner},
		{"COPY","./config/redis.conf /usr/local/etc/redis/redis.conf"},
		{"EXPOSE", app.Spec["port"]},
		{"ENTRYPOINT", "[\"redis-server\", \"/usr/local/etc/redis/redis.conf\"]"},
	}
	dockerfile:=generateDockerfile(app,properties)
	return []byte(dockerfile)
}

func (app *GenApp)generateDockerfile(){

	switch nature := app.App.Spec["nature"]; nature {
	case "staticApp":
		app.Dockerfile=GenerateApacheDockerfile(app.App)
	case "mongodb":
		app.Dockerfile=GenerateMongoDockerfile(app.App)
	case "mysql":
		app.Dockerfile=GenerateMysqlDockerfile(app.App)
	case "redis":
		app.Dockerfile=GenerateRedisDockerfile(app.App)
	case "TBD":
		fmt.Println("TBD.")
	default:
		fmt.Printf("NOT SUPPORTED")
	}
}

func (app *GenApp)pushDockerfile(ctx context.Context,user string)error  {
	dockerfileOptions := BuilFileOptions("Generating Dockerfile...", user, app.Dockerfile)
	_, err := app.CommitFile(ctx, "Dockerfile", dockerfileOptions)
	return err
}
