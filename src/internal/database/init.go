package database

import (
	"fmt"
	"database/sql"

    _"github.com/lib/pq"
)


//
func NewConnection(user, password, databaseName, host string) (*sql.DB, error){

	connStr := fmt.Sprintf("user=%s password=%s host=%s sslmode=disable", user, password, host);
	db, err := sql.Open("postgres", connStr);
	if err != nil {
		return nil ,err;
	}

	return db, nil
}