package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.Handle("/", http.HandlerFunc(home))
	http.HandleFunc("/dog/", http.HandlerFunc(dog))
	http.HandleFunc("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}

func home(resW http.ResponseWriter, req *http.Request) {
	tpl.Execute(resW, "Welcome!")
}

func dog(resW http.ResponseWriter, req *http.Request) {
	tpl.Execute(resW, "Woof woof!")
}

func me(resW http.ResponseWriter, req *http.Request) {
	tpl.Execute(resW, "Paolo")
}
