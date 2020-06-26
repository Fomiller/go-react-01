package main

import (
	"github.com/fomiller/go-react-01/server/API"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Default Gin router
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("../client/build", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", API.HomeHandler)
		api.GET("/pong", API.PongHandler)
	}

	router.Run(":5000")
}
