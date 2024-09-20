package main

import (
	"fmt"
	"net"
	"os"
)

var _ = net.Listen
var _ = os.Exit

func main() {

	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		fmt.Println("Failed to bind to port 9092")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	buffer := make([]byte, 1024)

	if _, err := conn.Read(buffer); err != nil {
		fmt.Println("conn Read error: ", err.Error())
	} else {
		fmt.Println("conn Read success")
	}

	if _, err := conn.Write([]byte{0, 0, 0, 0, 0, 0, 0, 7}); err != nil {
		fmt.Println("conn Write error: ", err.Error())
	} else {
		fmt.Println("conn Write success")
	}
}
