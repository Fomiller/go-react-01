package main

import (
	"net/http"

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
		api.GET("/", homeHandler)
		api.GET("/pong", pongHandler)
	}

	router.Run(":5000")
}

// Define Handlers
func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API Home",
	})
}

func pongHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PONG!",
	})
}
