package handlers

import (
	"log"
	"net/http"

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
	} else {

		dbusername, dbpassword := models.ReadUser(username, password)

		if username == dbusername && password == dbpassword {

			// http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

			log.Println("You have been logged in Successfully.")

		} else {
			log.Println("Invalid username or password!!")
		}
	}

	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}
