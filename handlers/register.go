package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register Handler
func RegisterHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}
