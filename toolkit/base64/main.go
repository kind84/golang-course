package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := `your string with characters not allowed:"`

	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(s64)
	fmt.Println(string(bs))
}
