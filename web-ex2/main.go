package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

type caliHotels struct {
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h1 := hotel{
		Name:    "Hilton",
		Address: "Beverly Hills",
		City:    "Los Angeles",
		Zip:     "29015",
		Region:  "Southern",
	}

	h2 := hotel{
		Name:    "Holiday Inn",
		Address: "Portobello",
		City:    "San Francisco",
		Zip:     "29011",
		Region:  "Central",
	}

	h3 := hotel{
		Name:    "Plaza",
		Address: "Main Road",
		City:    "Sacramento",
		Zip:     "20192",
		Region:  "Northern",
	}

	ch := caliHotels{[]hotel{h1, h2, h3}}

	err := tpl.Execute(os.Stdout, ch)
	if err != nil {
		log.Fatalln(err)
	}
}
