package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Logout Handler
func LogoutHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}
