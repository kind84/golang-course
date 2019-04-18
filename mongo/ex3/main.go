package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/kind84/golang-course/mongo/ex3/controllers"
	"github.com/kind84/golang-course/mongo/ex3/sessions"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	sessions.DbSessionsCleaned = time.Now()
}

func main() {
	gc := controllers.NewGeneralController(tpl)

	http.HandleFunc("/", gc.Index)
	http.HandleFunc("/bar", gc.Bar)
	http.HandleFunc("/signup", gc.Signup)
	http.HandleFunc("/login", gc.Login)
	http.HandleFunc("/logout", gc.Logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
