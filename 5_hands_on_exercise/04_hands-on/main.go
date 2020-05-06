package main

import (
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	defer conn.Close()

	respond(conn)
}

func respond(conn net.Conn) {
	io.WriteString(conn, "I am connected")
}
