package main

import (
	"html/template"
	"log"
	"net/http"
)

func hdl(resW http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Println(err)
	}

	tpl.Execute(resW, req.Form)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.HandleFunc("/", hdl)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
