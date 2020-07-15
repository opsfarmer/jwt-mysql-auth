package handlers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/parikshitg/jwt-mysql-auth/models"
)

// Register Handler
func GetRegister(c *gin.Context) {

	ok, _ := IsAuthenticated(c)
	if ok {
		location := url.URL{Path: "/welcome"}
		c.Redirect(http.StatusSeeOther, location.RequestURI())
		return
	}

	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}

func PostRegister(c *gin.Context) {

	// Reading form values
	username := c.PostForm("username")
	password := c.PostForm("password")
	password2 := c.PostForm("password2")

	// validating fields
	if username == "" || password == "" || password2 == "" {

		log.Println("Fields can not be empty!!")
		return
	}

	// Check if already registered
	dbusername, dbpassword := models.ReadUser(username, password)

	if dbusername != "" || dbpassword != "" {

		log.Println("User Already Registered!!")
		return
	}

	// If username already taken
	if username == dbusername {

		log.Println("This username is taken.")
		return
	}

	// If passwords didn't match
	if password != password2 {

		log.Println("Passwords Doesn't Match !!")
		return
	}

	// Create User
	models.CreateUser(username, password)
	log.Println("Registered Successfully.")

	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}
