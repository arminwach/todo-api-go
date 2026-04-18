package db

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite", "./todos.db")
	if err := DB.Ping(); err != nil {
		return err
	}
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(1)

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		task TEXT NOT NULL
	)`)
	if err != nil {
		return err
	}
	return nil
}
