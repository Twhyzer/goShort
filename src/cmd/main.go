package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/Twhyzer/goShort/src/api"
	"github.com/gorilla/mux"
)

// Initializes the router
func main() {
	fmt.Println("Starting the server on :80")

	r := mux.NewRouter()
	r.HandleFunc("/", api.HandleShortCreate()).Methods(http.MethodPost)
	r.HandleFunc("/", api.HandleShortDelete()).Methods(http.MethodDelete)
	r.HandleFunc("/{key}", api.HandleShortRedirect()).Methods(http.MethodGet)
	r.HandleFunc("/", resourceNotExistent)
	http.ListenAndServe(":80", r)
}

// If the requested URL does not exist, a 404 error should be returned.
func resourceNotExistent(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "404 Not Found (GoLang)")
}
