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

	var dbusername, dbid, dbcreatedat string

	ok, claims := IsAuthenticated(c)
	if !ok {
		location := url.URL{Path: "/login"}
		c.Redirect(http.StatusSeeOther, location.RequestURI())
		return
	}

	_, database := models.ExistingUser(claims.Username)

	if database == "test1" {

		query := "SELECT id, username, created_at  FROM test1table WHERE username = ?"
		if err := models.Db1.QueryRow(query, claims.Username).Scan(&dbid, &dbusername, &dbcreatedat); err != nil {
			log.Println("read user error : ", err)
		}
	} else {

		query := "SELECT id, username, created_at  FROM test2table WHERE username = ?"
		if err := models.Db2.QueryRow(query, claims.Username).Scan(&dbid, &dbusername, &dbcreatedat); err != nil {
			log.Println("read user error : ", err)
		}
	}

	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"title":      "Welcome Page",
		"message":    claims.Username,
		"id":         dbid,
		"username":   dbusername,
		"created_at": dbcreatedat,
		"present_in": database,
	})
}
