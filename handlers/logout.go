package handlers

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// Logout Handler
func LogoutHandler(c *gin.Context) {

	ok, _ := IsAuthenticated(c)
	if !ok {
		location := url.URL{Path: "/login"}
		c.Redirect(http.StatusSeeOther, location.RequestURI())
		return
	}

	c.SetCookie("auth_token", "", -1, "/", "localhost", false, true)

	location := url.URL{Path: "/login"}
	c.Redirect(http.StatusSeeOther, location.RequestURI())
}
