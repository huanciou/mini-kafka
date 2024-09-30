package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/*
	apiKey:
		1:
			- .messageLength
			- .correlation_id

		18:
		  - .messageLength
			- .correlation_id

	apiVersion:

		>=11:
			- .TAG_BUFFER

*/

type messageLengthStuct struct {
	MessageLength uint32
}

type correlationIdStruct struct {
	CorrelationId uint32
}

type tagBufferStruct struct {
	TagBuffer uint8
}

func HeaderResp(reqBuffer []byte) []byte {

	headerBuffer := new(bytes.Buffer)

	var (
		apiKey        uint16 = binary.BigEndian.Uint16(reqBuffer[4:6])
		apiVersion    uint16 = binary.BigEndian.Uint16(reqBuffer[6:8])
		correlationId uint32 = binary.BigEndian.Uint32(reqBuffer[8:12])
	)

	messageLengthChecker(headerBuffer)
	correlationIdChecker(headerBuffer, correlationId)
	tagBufferChecker(headerBuffer, apiKey, apiVersion)

	return headerBuffer.Bytes()
}

func messageLengthChecker(headerBuffer *bytes.Buffer) {
	field := messageLengthStuct{
		MessageLength: 0,
	}

	binary.Write(headerBuffer, binary.BigEndian, field)
}

func correlationIdChecker(headerBuffer *bytes.Buffer, correlationId uint32) {
	field := correlationIdStruct{
		CorrelationId: correlationId,
	}

	binary.Write(headerBuffer, binary.BigEndian, field)
}

func tagBufferChecker(headerBuffer *bytes.Buffer, apiKey, apiVersion uint16) {

	errorCode := ErrorCodeChecker(apiKey, apiVersion)

	isErrorVersion := errorCode != 0

	if apiVersion < 11 || isErrorVersion {
		return
	}

	fmt.Println("tag_buffer in header. version: ", apiVersion)

	field := tagBufferStruct{
		TagBuffer: 0,
	}

	binary.Write(headerBuffer, binary.BigEndian, field)
}
