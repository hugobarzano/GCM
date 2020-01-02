package handlers

import (
	"code-runner/internal/models"
	"errors"
	"net/http"
	"strings"
)

func getAppFromRequest(req *http.Request) (*models.App,error) {

	if err := req.ParseForm(); err != nil {
		return nil,err
	}

	appName:=req.FormValue("name")
	if appName==""{
		return nil,errors.New("name is mandatory")
	}

	appPort:=req.FormValue("port")
	if appPort==""{
		return nil,errors.New("port is mandatory")
	}

	appSpec:= make(map[string]string)
	appSpec["port"]=appPort
	appSpec["nature"]=req.FormValue("nature")

	app := &models.App{
		Name: strings.Replace(
			appName, "\"", "", -1),
		Des: strings.Replace(
			req.FormValue("description"), "\"", "", -1),
		Spec:appSpec,
	}
	return app,nil

}


