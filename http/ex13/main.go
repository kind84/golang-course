package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", handle)
	http.Handle("/public/pic/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handle(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
