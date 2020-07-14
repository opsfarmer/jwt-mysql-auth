package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"title": "Welcome Page",
	})
}