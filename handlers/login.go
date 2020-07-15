package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get Login Handler
func GetLogin(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

// Post Login Handler
func PostLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	log.Println("username : ", username, "password : ", password)

	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}
