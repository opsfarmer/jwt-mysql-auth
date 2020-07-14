package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home Handler
func HomeHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home Page",
	})
}
