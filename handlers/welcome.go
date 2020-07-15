package handlers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/parikshitg/jwt-mysql-auth/models"
)

// Welcome Handler
func WelcomeHandler(c *gin.Context) {

	ok, claims := IsAuthenticated(c)
	if !ok {
		location := url.URL{Path: "/login"}
		c.Redirect(http.StatusSeeOther, location.RequestURI())
		return
	}

	var dbusername, dbid, dbcreatedat string
	query := "SELECT id, username, created_at  FROM jwtusers WHERE username = ?"
	if err := models.Db.QueryRow(query, claims.Username).Scan(&dbid, &dbusername, &dbcreatedat); err != nil {
		log.Println("Read User Error : ", err)
	}

	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"title":      "Welcome Page",
		"message":    claims.Username,
		"id":         dbid,
		"username":   dbusername,
		"created_at": dbcreatedat,
	})
}
