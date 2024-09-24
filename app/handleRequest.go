package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func HandleRequest(conn net.Conn, buffer []byte) {
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

	numbersOfApiKeys := []byte{2}      // 1 Byte
	apiKey := buffer[4:6]              // 2 Bytes
	minVersion := []byte{0, 0}         // 2 Bytes
	maxVersion := []byte{0, 4}         // 2 Bytes
	tagFields := []byte{0}             // 1 Byte
	throttleTime := []byte{0, 0, 0, 0} // 4 Bytes

	/* 這邊 tester 要測試包含多個 Api key 版本, 需要 =2 */

	// numbersOfApiKeys := []byte{2}      // 1 Byte
	// apiKey := buffer[4:6]              // 2 Bytes
	// minVersion := []byte{0, 0}         // 2 Bytes
	// maxVersion := []byte{0, 4}         // 2 Bytes
	// tagFields := []byte{0}             // 1 Byte
	// throttleTime := []byte{0, 0, 0, 0} // 4 Bytes

	bodyResp := append(numbersOfApiKeys, apiKey...)
	bodyResp = append(bodyResp, minVersion...)
	bodyResp = append(bodyResp, maxVersion...)
	bodyResp = append(bodyResp, tagFields...)
	bodyResp = append(bodyResp, throttleTime...)
	bodyResp = append(bodyResp, tagFields...)

	if _, err := conn.Write(bodyResp); err != nil {
		fmt.Println("conn Write error: ", err.Error())
	}
}
