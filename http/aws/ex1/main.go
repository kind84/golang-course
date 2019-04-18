package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)

	http.ListenAndServe(":80", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello from AWS")
}
