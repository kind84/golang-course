package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ls, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer ls.Close()

	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	i := 0

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			url := strings.Fields(ln)[1]
			fmt.Println("***URL", url)
			switch url {
			case "/dog":
				respond(dog, conn)
			case "/cat":
				respond(cat, conn)
			case "/":
				respond(all, conn)
			}
		}
		if ln == "" {
			break
		}

		i++
	}
	defer conn.Close()

	fmt.Println("connection closed")
}

func dog(conn net.Conn) string {
	return `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>The Dog Page</title></head><body>
	<strong>Welcome to the Dog page</strong></body></html>`
}

func cat(conn net.Conn) string {
	return `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>The Cat Page</title></head><body>
	<strong>Welcome to the Cat page</strong></body></html>`
}

func all(conn net.Conn) string {
	return `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>The Home Page</title></head><body>
	<strong>Welcome to the Home page</strong></body></html>`
}

func respond(getBody func(conn net.Conn) string, conn net.Conn) {
	body := getBody(conn)

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
