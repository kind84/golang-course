package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

const key string = "amore"

func main() {
	msg := "Mary ti amo"
	fmt.Printf("message: \t%s\n", msg)
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(msg))
	hmsg := h.Sum(nil)
	fmt.Printf("hashed message: %x\n", hmsg)
	result := checkMAC([]byte("mary ti amo"), hmsg, []byte(key))
	fmt.Printf("match: \t\t%v\n", result)
}

func checkMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	fmt.Printf("expected hash: \t%x\n", expectedMAC)
	return hmac.Equal(messageMAC, expectedMAC)
}
