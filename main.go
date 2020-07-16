package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/parikshitg/jwt-mysql-auth/handlers"
	"github.com/parikshitg/jwt-mysql-auth/models"
)

func main() {

	log.SetFlags(log.Ltime | log.Lshortfile)

	r := gin.Default()

	r.LoadHTMLGlob("views/*")

	var err error
	// Open database test1
	models.Db1, err = sql.Open("mysql", "root:Poonam26#@tcp(127.0.0.1:3306)/test1")
	if err != nil {
		log.Println("Db Open Error:", err)
	}

	defer models.Db1.Close()
	log.Println("Successfully connected to Database. test1")

	// Open database test2
	models.Db2, err = sql.Open("mysql", "root:Poonam26#@tcp(127.0.0.1:3306)/test2")
	if err != nil {
		log.Println("Db Open Error:", err)
	}

	defer models.Db2.Close()
	log.Println("Successfully connected to Database. test2")

	// Open database test
	models.Db, err = sql.Open("mysql", "root:Poonam26#@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Println("Db Open Error:", err)
	}

	defer models.Db.Close()
	log.Println("Successfully connected to Database. test")

	// Create a database table in test1
	query1 := `
            CREATE TABLE IF NOT EXISTS test1table (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

	if _, err := models.Db1.Exec(query1); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully Created Table in test1")

	// Create a database table in test2
	query2 := `
            CREATE TABLE IF NOT EXISTS test2table (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

	if _, err := models.Db2.Exec(query2); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully Created Table in test2")

	// Create a database table in test2
	query := `
            CREATE TABLE IF NOT EXISTS testtable (
                username TEXT NOT NULL,
                dbname TEXT NOT NULL
            );`

	if _, err := models.Db.Exec(query); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully Created Table in test")

	// Routes
	r.GET("/", handlers.HomeHandler)
	r.GET("/login", handlers.GetLogin)
	r.POST("/login", handlers.PostLogin)
	r.GET("/logout", handlers.LogoutHandler)
	r.GET("/register", handlers.GetRegister)
	r.POST("/register", handlers.PostRegister)
	r.GET("/welcome", handlers.WelcomeHandler)

	// Static Files
	r.Static("/css", "static/css")
	r.Static("/js", "static/js")

	r.Run()
}
