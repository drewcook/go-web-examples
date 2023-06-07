package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id        int
	username  string
	password  string
	createdAt time.Time
}

func main() {
	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")
	checkErr(err)

	// Initialize the first connection to the database
	pingErr := db.Ping()
	checkErr(pingErr)

	CreateUsersTable(db)
	AddNewUser(db, "John", "Doe")
	AddNewUser(db, "Darth", "Vader")
	AddNewUser(db, "Leia", "Skywalker")
	AddNewUser(db, "Luke", "Skywalker")
	GetUser(db, 2)
	GetUsers(db)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateUsersTable(db *sql.DB) {
	// SQL query
	query := `
		CREATE TABLE users (
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
	);`
	// Executes the SQL query in our database
	_, err := db.Exec(query)
	checkErr(err)
}

func AddNewUser(db *sql.DB, name string, pw string) (int64, sql.Result) {
	query := `INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`
	// Inserts our data into the users table and returns with the result and a possible error.
	// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
	result, err := db.Exec(query, name, pw, time.Now())
	checkErr(err)
	// Get new ID
	userId, err := result.LastInsertId()
	checkErr(err)

	fmt.Printf("Added new user: %v\n", userId)

	return userId, result
}

func GetUser(db *sql.DB, userId int) User {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)
	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err := db.QueryRow(query, userId).Scan(&id, &username, &password, &createdAt)
	checkErr(err)

	fmt.Printf("Got user: %v\n", id)
	// transform data
	var u User
	return u
}

func GetUsers(db *sql.DB) []User {
	// query the rows, return back user structs
	query := `SELECT id, username, password, created_at FROM users`
	rows, queryErr := db.Query(query)
	checkErr(queryErr)
	defer rows.Close()
	// transform data
	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
		checkErr(err)
		users = append(users, u)
	}
	err := rows.Err()
	checkErr(err)

	fmt.Printf("Got all users: %#v\n", users)

	return users
}

func DeleteUser(db *sql.DB, userId int) int64 {
	query := `DELETE FROM users WHERE id = ?`
	res, err := db.Exec(query, userId)
	checkErr(err)
	id, err2 := res.LastInsertId()
	checkErr(err2)
	fmt.Printf("Deleted user: %v\n", id)
	return id
}
