package models

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db1 *sql.DB
var Db2 *sql.DB
var Db *sql.DB

// Create User in Test1 Database
func CreateUserTest1(username, password string) {

	createdAt := time.Now()

	_, err := Db1.Exec(`INSERT INTO test1table (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		log.Println("Insert Error : ", err)
	}

	_, err = Db.Exec(`INSERT INTO testtable (username, dbname) VALUES (?, ?)`, username, "test1")
	if err != nil {
		log.Println("Insert Error : ", err)
	}

}

// Create User in Test2 Database
func CreateUserTest2(username, password string) {

	createdAt := time.Now()

	_, err := Db2.Exec(`INSERT INTO test2table (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		log.Println("Insert Error : ", err)
	}

	_, err = Db.Exec(`INSERT INTO testtable (username, dbname) VALUES (?, ?)`, username, "test2")
	if err != nil {
		log.Println("Insert Error : ", err)
	}
}

// Checks if the user exists
func ExistingUser(username string) (bool, string) {

	var dbusername, dbdatabase string

	query := "SELECT username, dbname FROM testtable WHERE username = ?"
	if err := Db.QueryRow(query, username).Scan(&dbusername, &dbdatabase); err != nil {
		log.Println("Read User Error : ", err)
	}

	if dbusername == username {
		return true, dbdatabase
	}

	return false, ""
}

// Reads a user from test1 table
func ReadUserTest1(username, password string) (string, string) {

	var dbusername, dbpassword string

	query := "SELECT username, password FROM test1table WHERE username = ?"
	if err := Db1.QueryRow(query, username).Scan(&dbusername, &dbpassword); err != nil {
		log.Println("Read User Error : ", err)
	}

	return dbusername, dbpassword
}

// Reads a user from test2 table
func ReadUserTest2(username, password string) (string, string) {

	var dbusername, dbpassword string

	query := "SELECT username, password FROM test2table WHERE username = ?"
	if err := Db2.QueryRow(query, username).Scan(&dbusername, &dbpassword); err != nil {
		log.Println("Read User Error : ", err)
	}

	return dbusername, dbpassword
}
