package handlers

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

// Logout Handler
func LogoutHandler(c *gin.Context) {

	c.SetCookie("auth_token", "", 0, "/", "localhost", false, true)

	location := url.URL{Path: "/login"}
	c.Redirect(http.StatusSeeOther, location.RequestURI())
}
