package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li,err := net.Listen("tcp",":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for{
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handle(conn)
	}
}
func handle(conn net.Conn){
	defer conn.Close()

	request(conn)
}

func request(conn net.Conn){
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		ln := scanner.Text()
		fmt.Println(ln)
		if ln=="" {
			// End of header
			break
		}
	}
	fmt.Println("Headers Completed")
	io.WriteString(conn,"I got header")
}