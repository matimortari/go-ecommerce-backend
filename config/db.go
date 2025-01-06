package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Create a new PostgreSQL database connection
func NewPostgreSQLStorage(user, password, host, dbname string) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		user,
		password,
		host,
		dbname,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening PostgreSQL database: %v", err)
	}

	return db, nil
}
