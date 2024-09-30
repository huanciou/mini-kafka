package main

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/codecrafters-io/kafka-starter-go/app/protocol"
)

func HandleRequest(conn net.Conn, buffer []byte) {

	/*
		Request struct:
			header:
				buffer[0:4] -> message length / 4 BYTE
				buffer[4:6] -> request_api_key	INT16 / 2 BYTE
				buffer[6:8] -> request_api_version	INT16 / 2 BYTE
				buffer[8:12] -> correlation_id	INT32 / 4 BYTE
	*/

	var (
		apiKey     uint16 = binary.BigEndian.Uint16(buffer[4:6])
		apiVersion uint16 = binary.BigEndian.Uint16(buffer[6:8])
		bodyResp   []byte
	)

	headerResp := protocol.HeaderResp(buffer)

	switch apiKey {
	case 1: // Fetch API
		bodyResp = protocol.FetchAPIBodyResp(apiKey, apiVersion)
	case 18: // API Versions
		bodyResp = protocol.ApiVersionBodyResp(apiKey, apiVersion)
	default:
	}

	leng := uint32(len(bodyResp) + len(headerResp) - 4)
	binary.BigEndian.PutUint32(headerResp[:4], leng)

	if _, err := conn.Write(headerResp); err != nil {
		fmt.Println("Header Respond Error: ", err.Error())
	}

	if _, err := conn.Write(bodyResp); err != nil {
		fmt.Println("Body Respond Error: ", err.Error())
	}

}
