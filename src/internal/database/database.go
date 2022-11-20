package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DatabaseConnection struct {
	User         string
	DatabaseName string
	Host         string
	Database     *sql.DB
}

func NewConnection(user, password, databaseName, host string) (DatabaseConnection, error) {
	connStr := fmt.Sprintf("user=%s password=%s, dbname=%s, host=%s sslmode=disable", user, password, databaseName, host)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return DatabaseConnection{}, err
	}

	return DatabaseConnection{
		DatabaseName: databaseName,
		User:         user,
		Host:         host,
		Database:     db,
	}, nil
}
