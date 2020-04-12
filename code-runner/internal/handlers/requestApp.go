package handlers

import (
	"code-runner/internal/models"
	"net/http"
	"regexp"
	"strings"
)

type requestApp struct {
	App    *models.App
	Errors map[string]string
}

func getAppFromRequest(req *http.Request) (*requestApp, error) {

	if err := req.ParseForm(); err != nil {
		return nil, err
	}
	appName := req.FormValue("name")
	appPort := req.FormValue("port")
	appNature := req.FormValue("nature")
	appSpec := make(map[string]string)
	appSpec["port"] = appPort
	appSpec["nature"] = appNature

	app := &models.App{
		Name: strings.ToLower(
				strings.Replace(
					appName, "\"", "", -1)),
		Des: strings.Replace(
			req.FormValue("description"), "\"", "", -1),
		Spec: appSpec,
		Img:getImgFromNature(appNature),
	}
	reqApp := &requestApp{
		App:    app,
	}
	return reqApp, nil
}

func (appRequest *requestApp) validateRequest() bool {
	appRequest.Errors=make(map[string]string)

	if appRequest.App.Name == ""{
		appRequest.Errors["Name"] = "Name is mandatory. Lowercase."
	}

	if len(appRequest.App.Name) > 20{
		appRequest.Errors["Name"] = "Name too long. It has to be less than 20 chars."
	}

	re := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	if ok := re.Match([]byte(appRequest.App.Name)); !ok{
		appRequest.Errors["Name"] = "Name must contains only alphanumeric lowercase chars."
	}

	re = regexp.MustCompile("^[a-zA-Z0-9@_.,\\s\\w]*$")
	if ok := re.Match([]byte(appRequest.App.Des)); !ok{
		appRequest.Errors["Des"] = "Description must contains only alphanumeric chars, \"@\",\"_\",\".\" and \",\""
	}

	if len(appRequest.App.Des) > 512 {
		appRequest.Errors["Des"] = "Description too long. It has to be less than 512 chars."
	}


	if appRequest.App.Spec["port"] == "" {
		appRequest.Errors["Port"] = "Port is mandatory."
	}

	re = regexp.MustCompile("^[0-9]*$")
	if ok := re.Match([]byte(appRequest.App.Spec["port"])); !ok{
		appRequest.Errors["Port"] = "Port must contains only numeric chars."
	}

	if len(appRequest.App.Spec["port"]) > 7 || len(appRequest.App.Spec["port"]) < 2 {
		appRequest.Errors["Port"] = "Invalid por length. At least 2 digits and at most 7."
	}

	if appRequest.App.Spec["nature"] == "" {
		appRequest.Errors["Nature"] = "Nature is mandatory."
	}
	return len(appRequest.Errors) == 0
}

func getImgFromNature(nature string)string {
	switch nature {
	case "staticApp":
		return img
	case "dataService":
		return img
	case "api":
		return img
	case "devOps":
		return img
	default:
		return ""
	}
}