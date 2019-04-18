package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "check your terminal")
	fmt.Println(req.RequestURI)
}
