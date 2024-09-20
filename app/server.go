package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

var _ = net.Listen
var _ = os.Exit

func apiVersionCheck(apiVersion int) bool {
	if apiVersion >= 0 && apiVersion <= 4 {
		return true
	}
	return false
}

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

	buffer := make([]byte, 1024)

	if _, err := conn.Read(buffer); err != nil {
		fmt.Println("conn Read error: ", err.Error())
	} else {
		fmt.Println(string(buffer))
	}

	//	The request starts with a 4 byte length field
	//	buffer[0:4] -> request start string / 4 BYTE
	//	buffer[4:6] -> request_api_key	INT16 / 2 BYTE
	//	buffer[6:8] -> request_api_version	INT16 / 2 BYTE
	//	buffer[8:12] -> correlation_id	INT32 / 4 BYTE

	resp := make([]byte, 8)
	copy(resp[:4], []byte{0, 0, 0, 0})
	copy(resp[4:], buffer[8:12])

	apiVersion := binary.BigEndian.Uint16(buffer[6:8])
	apiVersionValidation := apiVersionCheck(int(apiVersion))

	if !apiVersionValidation {
		resp = append(resp, 0, 35)
	} else {
		resp = append(resp, 0, 0)
	}

	if _, err := conn.Write(resp); err != nil {
		fmt.Println("conn Write error: ", err.Error())
	} else {
		fmt.Println("conn Write success")
	}

	defer conn.Close()

}
