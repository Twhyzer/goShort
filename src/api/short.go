package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Twhyzer/goShort/src/internal/shorter"
	"github.com/Twhyzer/goShort/src/internal/util"
)

func HandlerFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		domain := r.FormValue("domain")

		shortKey, err := shorter.CreateShortURL(domain)
		shortDomain := util.CreateShortDomain(shortKey)


		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "Invalid Database Connection: ")
			io.WriteString(w, err.Error())
		}

		handleJSONResponse(w, shortDomain)
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
