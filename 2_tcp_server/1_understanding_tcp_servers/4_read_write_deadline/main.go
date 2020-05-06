package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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

		go handle(conn)
	}
}
func handle(conn net.Conn) {
	conn.SetDeadline(time.Now().Add(10 * time.Second))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "You said %s:", ln)
	}
	defer conn.Close()

	fmt.Println("10 Seconds Over")
}
