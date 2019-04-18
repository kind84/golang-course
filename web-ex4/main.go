package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"os"
)

type csvData struct {
	Date     string
	Open     string
	High     string
	Low      string
	Close    string
	Volume   string
	AdjClose string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	file, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.FieldsPerRecord = 7

	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	//csv := make([]csvData, len(records))

	// for i, rec := range records {
	// 	if i == 0 {
	// 		continue
	// 	}
	// 	csv = append(csv, csvData{rec[0], rec[1], rec[2], rec[3], rec[4], rec[5], rec[6]})
	// }

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, records)
}
