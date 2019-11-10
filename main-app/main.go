package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"main-app/internal/handlers"

)
func main() {

	router := gin.Default()

	// frontend
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// System Backend
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message": "pong",
			})
		})
	}

	// Start and run the server
	api.GET("/workspaces", handlers.GetWorkspace)
	api.POST("/workspace/join/:id", handlers.JoinWorkspace)
	router.Run(":3000")
}

