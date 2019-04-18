package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type httpStatus struct {
	Code        int    `json:"Code"`
	Description string `json:"Description"`
}

func main() {
	var data []httpStatus

	rcvd := `[{"Code":200,"Description":"StatusOK"},{"Code":301,"Description":"StatusMovedPermanently"},{"Code":302,"Description":"StatusFound"},{"Code":303,"Description":"StatusSeeOther"},{"Code":307,"Description":"StatusTemporaryRedirect"},{"Code":400,"Description":"StatusBadRequest"},{"Code":401,"Description":"StatusUnauthorized"},{"Code":402,"Description":"StatusPaymentRequired"},{"Code":403,"Description":"StatusForbidden"},{"Code":404,"Description":"StatusNotFound"},{"Code":405,"Description":"StatusMethodNotAllowed"},{"Code":418,"Description":"StatusTeapot"},{"Code":500,"Description":"StatusInternalServerError"}] `

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	for _, d := range data {
		fmt.Printf("%v: %v\n", d.Code, d.Description)
	}
}
