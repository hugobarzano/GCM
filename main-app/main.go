package main

import (
	"context"
	"github.com/auth0/go-jwt-middleware"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"main-app/internal/auth"
	"main-app/internal/database"
	"main-app/internal/handlers"
	"net/http"
)



var jwtMiddleWare *jwtmiddleware.JWTMiddleware



func main() {

	mongoClient := database.GetMongoClient("mongodb://cesarcor1:cesarcor1@ds049624.mlab.com:49624/cesarcorpdb")
	err := mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to MongoDB", err)
	}




	err = godotenv.Load()
	if err != nil {
		log.Print("Error loading .env file")
	}


	jwtMiddleWare = auth.InitAuthMiddleware()
	router := gin.Default()

	// frontend
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// backend
	api := router.Group("/back")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message": "pong",
			})
		})
	}

	// Start and run the server
	//api.GET("/workspace",auth.AuthMiddleware(jwtMiddleWare),handlers.MyWorkspace)
	api.GET("/workspaces",auth.AuthMiddleware(jwtMiddleWare),handlers.GetWorkspace)
	api.POST("/workspace/join/:id", auth.AuthMiddleware(jwtMiddleWare),handlers.JoinWorkspace)
	router.Run(":3000")
}

