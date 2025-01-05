package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	postgresMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/matimortari/go-backend/config"
	"github.com/matimortari/go-backend/db"
)

func main() {
	// Open a connection to PostgreSQL using the connection string
	db, err := db.NewPostgreSQLStorage(config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBHost, config.Envs.DBName)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// Initialize PostgreSQL migration driver
	driver, err := postgresMigrate.WithInstance(db, &postgresMigrate.Config{})
	if err != nil {
		log.Fatal("Error initializing migration driver: ", err)
	}

	// Initialize migration instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("Error creating migration instance: ", err)
	}

	// Check the current migration version
	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

	// Handle migration commands (up or down)
	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Error running migrations up: ", err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal("Error running migrations down: ", err)
		}
	}
}
