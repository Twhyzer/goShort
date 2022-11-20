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
	Database, err := database.NewConnection(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DATABASE_NAME"), os.Getenv("POSTGRES_HOST"))
	fmt.Printf("USER: %s", os.Getenv("POSTGRES_USER"))

	if err != nil {
		panic("Database connection failed")
	}

	_, err = database.InsertShortURL(Database.Database, "https://test.de", "https://t.de")

	if err != nil {
		fmt.Println("InsertShortURL failed", err)
	}

	fmt.Println("Starting the server on :80")

	r := mux.NewRouter()
	r.HandleFunc("/short", api.HandlerFunc(Database)).Methods(http.MethodPost)
	http.ListenAndServe(":80", r)

}
