package handlers

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// Welcome Handler
func WelcomeHandler(c *gin.Context) {

	ok, claims := IsAuthenticated(c)
	if !ok {
		location := url.URL{Path: "/login"}
		c.Redirect(http.StatusSeeOther, location.RequestURI())
		return
	}

	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"title":   "Welcome Page",
		"message": claims.Username,
	})
}
