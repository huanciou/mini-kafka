package main

import (
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		buffer := make([]byte, 1024)

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("conn Read error: ", err.Error())
			return // if client disconnected, get an io.EOF err. then return to defer func
		} else {
			fmt.Println("buffer content: ", buffer[:n])
		}

		go HandleRequest(conn, buffer[:n])
	}
}
