package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/parikshitg/jwt-mysql-auth/handlers"
	"github.com/parikshitg/jwt-mysql-auth/middlewares"
	"github.com/parikshitg/jwt-mysql-auth/models"
)

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	var err error
	// Open a database
	models.Db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/jwt")
	if err != nil {
		log.Println("Db Open Error:", err)
	}

	defer models.Db.Close()

	log.Println("Successfully connected to Database.")

	// Create a database table
	query := `
            CREATE TABLE IF NOT EXISTS jwtusers (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

	if _, err := models.Db.Exec(query); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully Created Table")

	// Routes
	r.GET("/", handlers.HomeHandler)
	r.GET("/login", middlewares.UnauthenticatedUser(), handlers.GetLogin)
	r.POST("/login", middlewares.UnauthenticatedUser(), handlers.PostLogin)
	r.GET("/logout", middlewares.AuthenticatedUser(), handlers.LogoutHandler)
	r.GET("/register", middlewares.UnauthenticatedUser(), handlers.GetRegister)
	r.POST("/register", middlewares.UnauthenticatedUser(), handlers.PostRegister)
	r.GET("/welcome", middlewares.AuthenticatedUser(), handlers.WelcomeHandler)

	// Static Files
	r.Static("/css", "static/css")
	r.Static("/js", "static/js")

	r.Run()
}
