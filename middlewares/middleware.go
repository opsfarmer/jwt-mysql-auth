package middlewares

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// Checks if user is authenticated (middleware)
func AuthenticatedUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		cookie, err := c.Cookie("auth_token")
		if err != nil {
			log.Println("Cookie error : ", err)
		}

		if cookie == "" {

			location := url.URL{Path: "/login"}
			c.Redirect(http.StatusSeeOther, location.RequestURI())
		}
	}
}

// checks unauthenticated user (middleware)
func UnauthenticatedUser() gin.HandlerFunc {

	return func(c *gin.Context) {

		cookie, err := c.Cookie("auth_token")
		if err != nil {
			log.Println("Cookie error : ", err)
		}

		if cookie != "" {

			location := url.URL{Path: "/welcome"}
			c.Redirect(http.StatusSeeOther, location.RequestURI())
		}
	}
}
