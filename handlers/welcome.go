package handlers

import (
	"github.com/gin-gonic/gin"
)

func WelcomeHandler(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "GET Welcome",
	})
}
