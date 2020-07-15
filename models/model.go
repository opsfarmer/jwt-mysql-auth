package models

import (
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
