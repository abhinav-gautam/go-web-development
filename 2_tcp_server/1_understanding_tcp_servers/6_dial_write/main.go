package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn,err := net.Dial("tcp","localhost:8080")

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	fmt.Fprint(conn,"I am a TCP Client")
}
