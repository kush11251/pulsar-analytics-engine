package utils

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDatabase() error {
	db, _ = sql.Open("sqlite3", "./pulsar.db")
	return db.Ping()
}
