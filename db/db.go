package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("postgres", "postgres://user:pass@localhost:5432/f1?sslmode=disable")
	return err
}
