package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func Client() {
	connection, errDial := net.Dial("tcp", "127.0.0.5:8080")

	clientReader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(connection)

	if errDial != nil {
		log.Fatal("ERROR: ", errDial)
		return
	}

	defer connection.Close()

	for {

		fmt.Print("Send your name: ")
		clientRequest, err := clientReader.ReadString('\n')

		if err != nil {
			log.Fatal("ERROR: ", err)
			return
		}

		fmt.Fprintf(connection, clientRequest+"\n")
		message, errMes := serverReader.ReadString('\n')

		if errMes != nil {
			log.Fatal("ERROR: ", errMes)
			return
		} else {
			fmt.Print("Message from server: " + "Hello, " + message)
			connection.Close()
			return
		}
	}
}
