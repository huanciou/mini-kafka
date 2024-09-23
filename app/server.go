package main

import (
	"encoding/binary"
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

	/*
		The request starts with a 4 byte length field:
			header:
				buffer[0:4] -> message length / 4 BYTE
				buffer[4:6] -> request_api_key	INT16 / 2 BYTE
				buffer[6:8] -> request_api_version	INT16 / 2 BYTE
				buffer[8:12] -> correlation_id	INT32 / 4 BYTE
			body:
	*/

	buffer := make([]byte, 1024)

	if _, err := conn.Read(buffer); err != nil {
		fmt.Println("conn Read error: ", err.Error())
	} else {
		fmt.Println(string(buffer))
	}

	/*
		The response order:
			header:
				buffer[0:4] -> message length / 4 Byte
				buffer[4:12] -> correlation_id / 4 Byte
				buffer[12:14] -> error code / 2 Byte
			body:
				#API_Keys -> Int_8 / 1 Byte
					API_Key -> Int_16 / 2 Byte
					min_version -> Int_16 / 2 Byte
					max_version -> Int_16 / 2 Byte
	*/

	leng := []byte{0, 0, 0, 19}   // 4 Bytes
	correlationId := buffer[8:12] // 4 Bytes
	var errorCode []byte          // 2 Bytes

	// fmt.Println(binary.BigEndian.Uint32(correlationId))

	version := binary.BigEndian.Uint16(buffer[6:8])

	switch version {
	case 0, 1, 2, 3, 4:
		errorCode = []byte{0, 0}
	default:
		errorCode = []byte{0, 35}
	}

	resp := append(leng, correlationId...)
	resp = append(resp, errorCode...)

	if _, err := conn.Write(resp); err != nil {
		fmt.Println("conn Write error: ", err.Error())
	}

	/* 這邊 tester 要測試包含多個 Api key 版本, 需要 =2 */

	numbersOfApiKeys := []byte{2}      // 1 Byte
	apiKey := buffer[4:6]              // 2 Bytes
	minVersion := []byte{0, 0}         // 2 Bytes
	maxVersion := []byte{0, 4}         // 2 Bytes
	tagFields := []byte{0}             // 1 Byte
	throttleTime := []byte{0, 0, 0, 0} // 4 Bytes

	bodyResp := append(numbersOfApiKeys, apiKey...)
	bodyResp = append(bodyResp, minVersion...)
	bodyResp = append(bodyResp, maxVersion...)
	bodyResp = append(bodyResp, tagFields...)
	bodyResp = append(bodyResp, throttleTime...)
	bodyResp = append(bodyResp, tagFields...)

	if _, err := conn.Write(bodyResp); err != nil {
		fmt.Println("conn Write error: ", err.Error())
	}

	fmt.Println(resp)
	fmt.Println(bodyResp)

	defer conn.Close()

}
