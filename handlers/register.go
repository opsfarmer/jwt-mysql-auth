package handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "GET Register",
	})
}
