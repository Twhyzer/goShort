package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Twhyzer/goShort/src/internal/database"
	"github.com/Twhyzer/goShort/src/internal/shorter"
)

func HandlerFunc(dbConn *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		domain := r.FormValue("domain")

		short_domain := shorter.CreateShortURL()

		_, err := database.InsertShortURL(dbConn, domain, short_domain)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "Invalid Database Connection: ")
			io.WriteString(w, err.Error())
		}

		handleJSONResponse(w, short_domain)
	}
}

func handleJSONResponse(w http.ResponseWriter, data string) {
	json, err := json.Marshal(data)

	if err != nil {
		io.WriteString(w, "JSON Creation Failed: ")
		io.WriteString(w, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
