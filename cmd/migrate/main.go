package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	postgresMigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/matimortari/go-ecom-backend/config"
)

// Main entry point for the migration tooling
func main() {
	db, err := config.NewPostgreSQLStorage(config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBHost, config.Envs.DBName)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	driver, err := postgresMigrate.WithInstance(db, &postgresMigrate.Config{})
	if err != nil {
		log.Fatal("Error initializing migration driver: ", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("Error creating migration instance: ", err)
	}

	v, d, _ := m.Version()
	log.Printf("Version: %d, dirty: %v", v, d)

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
