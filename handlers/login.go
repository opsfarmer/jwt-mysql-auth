package handlers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/parikshitg/jwt-mysql-auth/models"
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

	if username == "" || password == "" {

		log.Println("Fields can not be empty!!")
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Login",
		})
	} else {

		dbusername, dbpassword := models.ReadUser(username, password)

		if username == dbusername && password == dbpassword {

			location := url.URL{Path: "/welcome"}
			c.Redirect(http.StatusSeeOther, location.RequestURI())

			log.Println("You have been logged in Successfully.")

		} else {
			log.Println("Invalid username or password!!")
			c.HTML(http.StatusOK, "login.html", gin.H{
				"title": "Login",
			})
		}
	}
}
