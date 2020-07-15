package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register Handler
func GetRegister(c *gin.Context) {

	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}

func PostRegister(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	password2 := c.PostForm("password2")

	log.Println("Printing register post form values : ", username, password, password2)

	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}
