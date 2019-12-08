package deploy

import (
	"fmt"
	"testing"
	"github.com/CenturyLinkLabs/docker-reg-client/registry"
)

func Test_Hub(t *testing.T){

	c := registry.NewClient()
	auth, err := c.Hub.GetReadToken("ubuntu")


	imageID, err := c.Repository.GetImageID("ubuntu", "latest", auth)
	if err != nil {
		panic(err)
	}

	fmt.Println(imageID)
}