package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	mux := httprouter.New()
	mux.GET("/", jsoning)
	mux.POST("/go", golanging)

	http.ListenAndServe(":8080", mux)
}

func jsoning(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	p1 := person{
		First: "Paolo",
		Last:  "Grisoli",
		Age:   34,
	}

	p2 := person{
		First: "Maria",
		Last:  "Nesca",
		Age:   26,
	}

	ppl := getPpl(p1, p2)

	err := json.NewEncoder(w).Encode(ppl)
	if err != nil {
		log.Println(err)
	}
}

func golanging(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var ppl []person

	err := json.NewDecoder(req.Body).Decode(&ppl)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(ppl)
}

func getPpl(ppl ...person) []person {
	return ppl
}
