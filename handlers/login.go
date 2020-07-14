package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login Handler
func LoginHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}
