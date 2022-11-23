package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Twhyzer/goShort/src/internal/database"
	"github.com/Twhyzer/goShort/src/internal/shorter"
	"github.com/Twhyzer/goShort/src/internal/util"
)

// Controls the process for creating a new short.
func HandleShortCreate() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20)
		domain := r.PostForm.Get("domain")

		requestKey, err := shorter.CreateShortURL(domain)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "Invalid Database Connection: ")
			io.WriteString(w, err.Error())
			return
		}

		shortDomain := util.CreateShortDomain(requestKey)

		handleJSONResponse(w, shortDomain)
	}
}


func HandleShortDelete() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20)
		key := r.PostForm.Get("key")

		err := database.DeleteShortByKey(key)

		if err != nil {
			handleJSONResponse(w, "The key was not found.")
		}

		handleJSONResponse(w, "Success")
	}
}

// Controls the process for creating a new short.
func HandleShortRedirect() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path[1:]

		short, err := database.GetShortByKey(key)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			io.WriteString(w, "{\"error\": \"Invalid Key\"}")
			return
		}

		http.Redirect(w, r, short.TargetUrl, http.StatusPermanentRedirect)
	}
}

// Writes the response of an HTTP request.
func handleJSONResponse(w http.ResponseWriter, data string) {
	json, err := json.Marshal(data)

	if err != nil {
		io.WriteString(w, "JSON Creation Failed: ")
		io.WriteString(w, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
