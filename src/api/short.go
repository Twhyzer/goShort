package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Twhyzer/goShort/src/internal/database"
)

func HandlerFunc(Connection database.DatabaseConnection) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := database.InsertShortURL(Connection.Database, "https://test.de", "https://t.de")

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "Invalid Database Connection: ")
			io.WriteString(w, err.Error())
		}

		shortsByte, err := json.Marshal("https://t.de")

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "JSON Creation Failed: ")
			io.WriteString(w, err.Error())
		}

		w.Write(shortsByte)
	}
}
