package main

import (
	"fmt"
	"net/http"

	"github.com/Twhyzer/goShort/src/api"
	"github.com/Twhyzer/goShort/src/internal/database"

	"github.com/gorilla/mux"
)

func main() {
	dbConn, dbConnErr := database.NewAppDatabase()

	if dbConnErr != nil {
		panic(dbConnErr.Error())
	}

	fmt.Println("Starting the server on :80")

	r := mux.NewRouter()
	r.HandleFunc("/short", api.HandlerFunc(dbConn)).Methods(http.MethodPost)
	http.ListenAndServe(":80", r)

}
