package handlers

import (
	"database/sql"
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
		log.Println("read user error : ", err)
	}

	// Reading Data From User Specific Database (starts)
	db, err := sql.Open("mysql", "root:Poonam26#@tcp(127.0.0.1:3306)/"+claims.Username)
	if err != nil {
		log.Println("user specific db open error:", err)
		return
	}

	var dbuname, dbabout string
	query2 := "SELECT username, about FROM details WHERE username = ?"
	if err := db.QueryRow(query2, claims.Username).Scan(&dbuname, &dbabout); err != nil {
		log.Println("read user details error : ", err)
		return
	}

	defer db.Close()
	// Reading Data From User Specific Database (ends)

	c.HTML(http.StatusOK, "welcome.html", gin.H{
		"title":      "Welcome Page",
		"message":    claims.Username,
		"id":         dbid,
		"username":   dbusername,
		"created_at": dbcreatedat,
		"about":      dbabout,
		"dbuname":    dbuname,
	})
}
