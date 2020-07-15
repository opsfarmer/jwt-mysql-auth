package handlers

import (
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/parikshitg/jwt-mysql-auth/models"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string
	jwt.StandardClaims
}

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

			// Jwt Starts--------------------------------------------------------------------------------------

			// tokens expiration time
			expirationTime := time.Now().Add(5 * time.Minute)

			claims := &Claims{
				Username: username,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtKey)
			if err != nil {
				log.Println(http.StatusInternalServerError)
				return
			}

			// set Cookie
			c.SetCookie("auth_token", tokenString, 3600, "/", "localhost", false, true)

			// Jwt Ends------------------------------------------------------------------------------------------

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
