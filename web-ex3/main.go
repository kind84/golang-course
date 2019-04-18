package main

import (
	"log"
	"os"
	"text/template"
)

type restaurant struct {
	Name string
	Menu menu
}

type menu struct {
	Breakfast []string
	Lunch     []string
	Dinner    []string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	res1 := restaurant{
		Name: "Road House",
		Menu: menu{
			Breakfast: []string{"eggs", "bacon", "fruit jam", "orange juice", "peanut butter", "toasted bread"},
			Lunch:     []string{"spaghetti", "t-bone steak", "salad", "grilled vegetables"},
			Dinner:    []string{"rice & curry", "spicy beans", "broccoli cream"},
		},
	}

	res2 := restaurant{
		Name: "Mexican Tapas",
		Menu: menu{
			Breakfast: []string{"milk", "chiken breast", "fruit jam", "orange juice"},
			Lunch:     []string{"fajitas", "burritos", "chicken wings", "grilled vegetables"},
			Dinner:    []string{"tacos", "spicy beans", "broccoli cream"},
		},
	}

	rests := struct {
		Restaurants []restaurant
	}{
		Restaurants: []restaurant{res1, res2},
	}

	err := tpl.Execute(os.Stdout, rests)
	if err != nil {
		log.Fatalln(err)
	}
}
