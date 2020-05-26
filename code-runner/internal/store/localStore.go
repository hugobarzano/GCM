package store

import (
	"code-runner/internal/constants"
	"encoding/base64"
	"io/ioutil"
	"log"
)

var SinglePageImg string
var ApiRestImg string
var DataServiceImg string
var DevOpsServiceImg string

func LoadImg(path string) string {
	imageFile, err := ioutil.ReadFile(path)

	if err != nil {
		log.Println(err.Error())
		return ""
	}
	encodedString := base64.StdEncoding.EncodeToString(imageFile)
	return encodedString
}
func setupImg(nature string) string {
	switch nature {
	case constants.SinglePage:
		return SinglePageImg
	case constants.ApiRest:
		return ApiRestImg
	case constants.DataService:
		return DataServiceImg
	case constants.DevOpsService:
		return DevOpsServiceImg
	default:
		return ""
	}
}
