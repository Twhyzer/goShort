package database

import (
	"database/sql"
)

type Shorts struct {
	Id         int
	TargetUrl  string
	RequestKey string
	Redirects  int
	Timestamp  string
}

// Query to add a new short to the database
func InsertShortURL(targetUrl string, shorturl string) (string, error) {
	dbConn := createDatabaseConnection()
	stmt, err := dbConn.Prepare("INSERT INTO links (targetUrl, requestKey) VALUES ($1, $2)")

	if err != nil {
		return "Prepare Statement Error", err
	}

	_, err = stmt.Exec(targetUrl, shorturl)

	if err != nil {
		return "Execute Error", err
	}

	return shorturl, nil
}

// Query to get short with key
func GetShortByKey(requestKey string) (Shorts, error) {
	return getShortBy("SELECT * FROM links WHERE requestKey = $1", requestKey)
}

// Query to delete short with key
func DeleteShortByKey(key string) error {
	dbConn := createDatabaseConnection()
	stmt, err := dbConn.Prepare("DELETE FROM links WHERE requestKey = $1")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(key)

	if err != nil {
		return err
	}

	return nil
}

// Returns a short according to the query
func getShortBy(query string, value string) (Shorts, error) {
	dbConn := createDatabaseConnection()
	stmt, err := dbConn.Prepare(query)

	if err != nil {
		return Shorts{}, err
	}

	var short Shorts

	err = stmt.QueryRow(value).Scan(&short.Id, &short.TargetUrl, &short.RequestKey, &short.Redirects, &short.Timestamp)

	if err != nil {
		return short, err
	}

	return short, nil
}

// Returns a connection to the database
func createDatabaseConnection() *sql.DB {
	dbConn, dbConnErr := NewAppDatabase()

	if dbConnErr != nil {
		panic(dbConnErr.Error())
	}

	return dbConn
}
