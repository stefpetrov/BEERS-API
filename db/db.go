package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createBeersTable := `
	CREATE TABLE IF NOT EXISTS beers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price TEXT NOT NULL,
		image TEXT NOT NULL,
		rating_average REAL NOT NULL,
		rating_reviews INTEGER NOT NULL
	)
	`
	_, err := DB.Exec(createBeersTable)

	if err != nil {
		panic("Could not create beers table." + err.Error())
	}

}
