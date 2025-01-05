package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/matimortari/go-backend/cmd/api"
	"github.com/matimortari/go-backend/config"
)

func main() {
	// Define the PostgreSQL connection string using DBHost and DBPort
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Envs.DBUser,
		config.Envs.DBPassword,
		config.Envs.DBHost,
		config.Envs.DBPort,
		config.Envs.DBName,
	)

	// Initialize the PostgreSQL database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	// Test the database connection
	initStorage(db)

	// Start the API server
	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	log.Println("Connected to PostgreSQL database")
}
