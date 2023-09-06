package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error connecting...")
		os.Exit(1)
	}
	defer connection.Close()

	_, err = connection.Write([]byte("Hello server!"))

	//read message from server
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)

	if err != nil {
		fmt.Println("Error read...")
		os.Exit(1)
	}

	fmt.Println("Received data: ", string(buffer[:mLen]))
}
