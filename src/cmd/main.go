package main

import (
    "fmt"
    "net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, "<h1>Hello World!</h1>")
}

func main() {
    mux := &http.ServeMux{}
    mux.HandleFunc("/", handlerFunc)
    fmt.Println("Starting the server on :80")
    http.ListenAndServe(":80", mux)
}