package db

import (
	"database/sql"
	"log"
)

func ConectDB(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	db.Exec("PRAGMA busy_timeout = 5000;")

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT,
		age INTEGER
	);`

	_, err := db.Exec(query)
	return err
}
