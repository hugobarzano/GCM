package models

import (
	"fmt"
	"testing"
)

func Test_GetImagenName(t *testing.T) {
	app := &App{
		Name:  "appName",
		Owner: "appOwner",
	}

	img := app.GetImageName()
	log.Println(img)
}
