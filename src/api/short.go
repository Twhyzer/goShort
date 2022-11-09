package api

import (
	"encoding/json"
	"io"
	"net/http"
)

func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	shortsByte, err := json.Marshal("hello world!")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "JSON Creation Failed")
		io.WriteString(w, err.Error())
	}

	w.Write(shortsByte)
}
