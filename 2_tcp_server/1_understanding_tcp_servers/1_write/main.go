package main

import (
	"fmt"
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
			log.Println(err)
			continue
		}

		io.WriteString(conn, "\nHello from TCP Server\n")
		fmt.Fprint(conn, "\rWhat is your name?\n")
		fmt.Fprintf(conn, "%v", "\rIs it Abhinav?")

		conn.Close()
	}
}
