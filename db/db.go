package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// NewPostgreSQLStorage initializes a new PostgreSQL database connection.
func NewPostgreSQLStorage(user, password, host, dbname string) (*sql.DB, error) {
	// Build the connection string for PostgreSQL
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		user,
		password,
		host,
		dbname,
	)

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening PostgreSQL database: %v", err)
	}

	// Return the database connection
	return db, nil
}
