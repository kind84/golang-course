package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}

func home(resW http.ResponseWriter, req *http.Request) {
	io.WriteString(resW, "Welcome!")
}

func dog(resW http.ResponseWriter, req *http.Request) {
	io.WriteString(resW, "Woof woof!")
}

func me(resW http.ResponseWriter, req *http.Request) {
	io.WriteString(resW, "Paolo")
}
