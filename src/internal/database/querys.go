package database

import (
	"database/sql"
)


func InsertShortURL(db *sql.DB, inputurl string, shorturl string) (string, error){
	stmt, err := db.Prepare("INSERT INTO links (inputurl, shorturl) VALUES ($1, $2)"); 

	if err != nil {
		return "Prepare Statement Error", err;
	}

	_, err = stmt.Exec(inputurl, shorturl)

	if err != nil {
		return "Execute Error", err
	}

	return "", nil
}