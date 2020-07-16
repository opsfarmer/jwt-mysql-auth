package handlers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/parikshitg/jwt-mysql-auth/models"
)

// Register Get Handler
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

// Register Post Handler
func PostRegister(c *gin.Context) {

	// Logging for Debugging
	log.SetFlags(log.Ltime | log.Lshortfile)

	// Reading form values
	username := c.PostForm("username")
	password := c.PostForm("password")
	password2 := c.PostForm("password2")
	database := c.PostForm("database")

	var flash string

	// validating fields
	if username == "" || password == "" || password2 == "" || database == "" {

		flash = "Fields can not be empty!!"
		log.Println(flash)
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register",
			"flash": flash,
		})
		return
	}

	// Check if already registered
	exists, _ := models.ExistingUser(username)
	if exists {

		flash = "User Already Registered!!"
		log.Println(flash)
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register",
			"flash": flash,
		})

		return
	}

	// If passwords didn't match
	if password != password2 {

		flash = "Passwords Doesn't Match !!"
		log.Println(flash)
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register",
			"flash": flash,
		})
		return
	}

	if database == "test1" {

		models.CreateUserTest1(username, password)
		flash = "Registered Successfully. Please Login."
		log.Println(flash)
	} else {

		models.CreateUserTest2(username, password)
		flash = "Registered Successfully. Please Login."
		log.Println(flash)
	}

	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
		"flash": flash,
	})
}
