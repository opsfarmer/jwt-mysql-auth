package models

import (
	"fmt"
	"log"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

type User struct {
	Id        string
	Username  string
	Password  string
	CreatedAt string
}

// Create user function creates a new user in  table
func CreateUser(username, password string) {

	createdAt := time.Now()

	_, err := Db.Exec(`INSERT INTO jwtusers (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		log.Println("Insert Error : ", err)
	}
}

// Read user reads a user from  table
func ReadUser(username, password string) (string, string) {

	var dbusername, dbpassword string

	query := "SELECT username, password FROM jwtusers WHERE username = ?"
	if err := Db.QueryRow(query, username).Scan(&dbusername, &dbpassword); err != nil {
		log.Println("Read User Error : ", err)
	}

	return dbusername, dbpassword
}

// Creating User Specific Database
func CreateUserDatabase(username string) {

	// Database connection
	db, err := sql.Open("mysql", "root:Poonam26#@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database created successfully")
	}

	// Create db
	_, err = db.Exec("CREATE DATABASE " + username)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database")
	}

	// Use datbase
	_, err = db.Exec("USE " + username)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("DB selected successfully..")
	}

	// Create table
	query := `
            CREATE TABLE IF NOT EXISTS details (
                username TEXT NOT NULL,
                about TEXT NOT NULL
            );`

	if _, err := db.Exec(query); err != nil {
		log.Println("details table creating error.", err)
	}

	log.Println("Created table successfully.")

	// Insert
	about := username + " is a good human being." // Creating a custom string for example
	_, err = db.Exec("INSERT INTO details (username, about) VALUES (?, ?)", username, about)
	if err != nil {
		log.Println("insert Error : ", err)
	}

	defer db.Close()
}
