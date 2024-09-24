package main

import (
	"fmt"
	"net"
)

func HandleConnection(conn net.Conn) {

	defer conn.Close()

	/*
		The request starts with a 4 byte length field:
			header:
				buffer[0:4] -> message length / 4 BYTE
				buffer[4:6] -> request_api_key	INT16 / 2 BYTE
				buffer[6:8] -> request_api_version	INT16 / 2 BYTE
				buffer[8:12] -> correlation_id	INT32 / 4 BYTE
			body:
	*/

	for {
		buffer := make([]byte, 1024)

		if _, err := conn.Read(buffer); err != nil {
			fmt.Println("conn Read error: ", err.Error())
			return // if client disconnected, get an io.EOF err. then return to defer func
		}

		go HandleRequest(conn, buffer)
	}
}
