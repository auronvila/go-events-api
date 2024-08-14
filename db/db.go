package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB
var err error

func InitDb() {
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic(fmt.Sprintf("Could not initialize the database: %v", err))
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
)
`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(fmt.Sprintf("Error in creating the users table: %v", err))
	}

	createEventsTable := `
CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME NOT NULL,
	userId INTEGER,
	FOREIGN KEY(userId) REFERENCES users(id)
)
`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic(fmt.Sprintf("Error in creating the events table: %v", err))
	}

	createRegistrationsTable := `
CREATE TABLE IF NOT EXISTS registrations (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    event_id INTEGER,
    user_id INTEGER,
    FOREIGN KEY(event_id) REFERENCES events(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
)
`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		panic(fmt.Sprintf("Error in creating the registrations table: %v", err))
	}
}
