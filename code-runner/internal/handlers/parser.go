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

	appSpec:= make(map[string]string)
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


