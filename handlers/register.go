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

	var flash string

	// validating fields
	if username == "" || password == "" || password2 == "" {

		flash = "Fields can not be empty!!"
		log.Println(flash)
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register",
			"flash": flash,
		})
		return
	}

	// Check if already registered
	dbusername, dbpassword := models.ReadUser(username, password)

	if dbusername != "" || dbpassword != "" {

		flash = "User Already Registered!!"
		log.Println(flash)
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register",
			"flash": flash,
		})
		return
	}

	// If username already taken
	if username == dbusername {

		flash = "This username is taken."
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

	// Create User
	models.CreateUser(username, password)
	flash = "Registered Successfully. Please Login."
	log.Println(flash)

	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
		"flash": flash,
	})
}
