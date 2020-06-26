package API

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define Handlers
func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "API Home",
	})
}

func PongHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "PONG!",
	})
}
