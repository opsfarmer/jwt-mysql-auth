package handlers

import (
	"github.com/gin-gonic/gin"
)

// Home Handler
func HomeHandler(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "GET Home",
	})

	// c.HTML(http.StatusOK, "homepage.html", gin.H{
	// 	"title": "Home Page",
	// })
}
