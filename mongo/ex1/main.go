package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/kind84/golang-course/mongo/ex1/controllers"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
