package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Welcome Handler
func WelcomeHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"title": "Welcome Page",
	})
}
