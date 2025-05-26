package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {

	connStr := "host=localhost port=5432 user=postgres password=123456 dbname=event-booking sslmode=disable"
	var err error

	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	err = DB.Ping()
	if err != nil {
		panic("Failed to ping the database: " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()

	fmt.Println("Successfully connected to PostgreSQL!")
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Failed to create users table: " + err.Error())
	}
}
