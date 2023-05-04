package main

import (
	"fmt"
	"log"

	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const (
	// these are constants because the db credentials should not change
	HOST     = "localhost"
	PORT     = 5432
	USER     = "root"
	PASSWORD = "secret"
	DBNAME   = "jwt"
)

func main() {
	// make db connection
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)

	runMigrateUp(connString)
}

func runMigrateUp(connString string) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("error when trying to connect:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("db instance is nil:", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://sql/migrations", "postgres", driver)
	if err != nil {
		log.Fatal("migration files not found:", err)
	}

	// migrate up
	upErr := m.Up()
	if upErr != nil {
		if upErr == migrate.ErrNoChange {
			log.Println("migration is up to date:", upErr)
		} else {
			log.Fatal("migration up error:", upErr)
		}
	}

	if upErr == nil {
		log.Println("database migrated up successfully")
	}
}
