package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Connects to the database and migrates it if necessary.
func NewDatabase() (*sql.DB, error) {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DATABASE_NAME"),
	)

	db, databaseConnectionErr := sql.Open("postgres", psqlconn)
	if databaseConnectionErr != nil {
		log.Fatalf("Could not Connect to Database! %s\n", databaseConnectionErr.Error())
		os.Exit(1)
	}

	MigrateDatabase(db, os.Getenv("POSTGRES_DATABASE_NAME"))

	return db, nil
}

// Returns a connection to the database
func CreateDatabaseConnection() *sql.DB {
	dbConn, dbConnErr := NewDatabase()

	if dbConnErr != nil {
		panic(dbConnErr.Error())
	}

	return dbConn
}
