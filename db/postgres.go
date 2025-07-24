package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := "host=localhost port=5432 user=postgres password=12345678 dbname=userservice sslmode=disable"
	return sql.Open("postgres", connStr)
}
