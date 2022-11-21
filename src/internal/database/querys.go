package database

import "database/sql"


func InsertShortURL(inputurl string, shorturl string) (string, error){
	dbConn := createDatabaseConnection()
	stmt, err := dbConn.Prepare("INSERT INTO links (inputurl, shorturl) VALUES ($1, $2)"); 

	if err != nil {
		return "Prepare Statement Error", err;
	}

	_, err = stmt.Exec(inputurl, shorturl)

	if err != nil {
		return "Execute Error", err
	}

	return shorturl, nil
}

func createDatabaseConnection() *sql.DB {
	dbConn, dbConnErr := NewAppDatabase()

	if dbConnErr != nil {
		panic(dbConnErr.Error())
	}

	return dbConn
}