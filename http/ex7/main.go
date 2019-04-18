package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)

	var reqLn []string

	for scanner.Scan() {
		ln := scanner.Text()
		reqLn = strings.Fields(ln)
		meth := reqLn[0]
		url := reqLn[1]
		if meth == "GET" && url == "/" {
			home(conn)
		}
		if meth == "GET" && url == "/apply" {
			apply(conn)
		}
		if meth == "POST" && url == "/apply" {
			applyPost(conn)
		}
		break
	}
}

func home(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Hello</title></head>
	<body><h1>Hello</h1></body></html>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Hello</title></head>
	<body><h1>Apply</h1><form action="/apply" method="POST"><input type="submit" value="Apply"></form></body></html>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func applyPost(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Hello</title></head>
	<body><h1>Apply Post</h1></body></html>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
