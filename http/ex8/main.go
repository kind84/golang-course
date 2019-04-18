package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	mux := httprouter.New()
	mux.GET("/", home)
	mux.GET("/dog", dog)
	mux.GET("/me", me)

	http.ListenAndServe(":8080", mux)
}

func home(resW http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	io.WriteString(resW, "Welcome!")
}

func dog(resW http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	io.WriteString(resW, "Woof woof!")
}

func me(resW http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	io.WriteString(resW, "Paolo")
}
