package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func main() {
	log = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: new(logrus.TextFormatter),
		Level:     logrus.DebugLevel,
	}

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Error(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(30 * time.Second))
	if err != nil {
		log.Info("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	log.Debug("***CODE GOT HERE***")
}
