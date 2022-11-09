package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Twhyzer/goShort/src/api"
	"github.com/Twhyzer/goShort/src/internal/database"

	"github.com/gorilla/mux"
)


func main() {
	_, err := database.NewConnection(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DATABASE_NAME"), os.Getenv("POSTGRES_HOST"))
	
	if err != nil {
		panic("Database connection failed")
	}

	fmt.Println("Starting the server on :80")

	r := mux.NewRouter()
	r.HandleFunc("/short", api.HandlerFunc).Methods(http.MethodPost)
	http.ListenAndServe(":80", r)

}