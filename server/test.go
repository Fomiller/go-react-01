package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func testHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Test working!",
	})
}
