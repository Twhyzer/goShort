package main

import (
	"fmt"
	"net/http"

	"github.com/Twhyzer/goShort/src/api"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting the server on :80")

	r := mux.NewRouter()
	r.HandleFunc("/short", api.HandlerFunc()).Methods(http.MethodPost)
	http.ListenAndServe(":80", r)

}
