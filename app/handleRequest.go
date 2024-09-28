package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func HandleRequest(conn net.Conn, buffer []byte) {

	fmt.Println("buffer leng: ", len(buffer))
	fmt.Println("buffer content: ", buffer)
	/*
		Request struct:
			header:
				buffer[0:4] -> message length / 4 BYTE
				buffer[4:6] -> request_api_key	INT16 / 2 BYTE
				buffer[6:8] -> request_api_version	INT16 / 2 BYTE
				buffer[8:12] -> correlation_id	INT32 / 4 BYTE
	*/

	/*
		Response struct:
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

	fmt.Println("handling request here ")

	correlationId := buffer[8:12] // 4 Bytes
	var errorCode []byte          // 2 Bytes

	version := binary.BigEndian.Uint16(buffer[6:8])

	switch version {
	case 0, 1, 2, 3, 4:
		errorCode = []byte{0, 0}
	default:
		errorCode = []byte{0, 35}
	}

	apiKey := binary.BigEndian.Uint16(buffer[4:6])

	var bodyResp []byte

	switch int(apiKey) {
	case 18: // APIVersions
		bodyResp = APIVersions(apiKey)
	default:
		bodyResp = []byte{}
	}

	fmt.Println(len(bodyResp))

	leng := []byte{0, 0, 0, byte(len(bodyResp) + 6)} // 4 Bytes

	fmt.Println(leng)

	HeaderResp := append(leng, correlationId...)
	HeaderResp = append(HeaderResp, errorCode...)

	fmt.Println("header", HeaderResp)

	// if _, err := conn.Write(HeaderResp); err != nil {
	// 	fmt.Println("Header Respond Error: ", err.Error())
	// }

	// if _, err := conn.Write(bodyResp); err != nil {
	// 	fmt.Println("Body Respond Error: ", err.Error())
	// }

	resp := append(HeaderResp, bodyResp...)

	if _, err := conn.Write(resp); err != nil {
		fmt.Println("Body Respond Error: ", err.Error())
	}

	fmt.Println(resp)
	// [0 0 0 19 84 246 20 22 0 0 2 0 18 0 0 0 4 0 0 0 0 0 0 0]

	fmt.Println("req done")
}
