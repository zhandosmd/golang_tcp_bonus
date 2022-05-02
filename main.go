package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("ERROR: ", err)
		return
	}

	fmt.Println("Server is starting...")

	go Client()

	conn, err := ln.Accept()

	if err != nil {
		log.Fatal("ERROR: ", err)
		return
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			log.Fatal("ERROR: ", err)
		}

		conn.Write([]byte(message + "\n"))
	}
}
