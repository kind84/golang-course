package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handle)
	http.HandleFunc("/favicon.ico", http.NotFound)

	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {

	c, err := req.Cookie("my-cookie")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	c.Value = strconv.Itoa(count)

	http.SetCookie(w, c)

	fmt.Fprintf(w, "This is your visit number %s", c.Value)
}
