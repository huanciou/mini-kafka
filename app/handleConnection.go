package main

import (
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()

	hC := 1
	REQ := 1

	for {
		buffer := make([]byte, 1024)

		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("conn Read error: ", err.Error())
			return // if client disconnected, get an io.EOF err. then return to defer func
		} else {
			fmt.Println("get req:", REQ)
			REQ++
		}

		fmt.Println("handle Connection Count: ", hC)
		go HandleRequest(conn, buffer[:n])
		hC++
	}
}
