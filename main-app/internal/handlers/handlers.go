package handlers

import (
	"main-app/internal/models"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

var workspaces = []models.Workspace{
	models.Workspace{1, 0, "this is a workspace to generate apps"},
	models.Workspace{2, 0, "bullshit"},
}



func GetWorkspace(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, workspaces)
}

func JoinWorkspace(c *gin.Context) {
	if wsId, err := strconv.Atoi(c.Param("id")); err == nil {
		for i := 0; i < len(workspaces); i++ {
			if workspaces[i].ID == wsId {
				workspaces[i].Members += 1
			}
		}
		c.JSON(http.StatusOK, &workspaces)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
