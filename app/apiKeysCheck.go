package main

import (
	"bytes"
	"encoding/binary"
)

/*
	numbersOfApiKeys := []byte{}      // 1 Byte
		apiKey := buffer[4:6]              // 2 Bytes
		minVersion := []byte{0, 0}         // 2 Bytes
		maxVersion := []byte{0, 4}         // 2 Bytes
		tagFields := []byte{0}             // 1 Byte
	throttleTime := []byte{0, 0, 0, 0} // 4 Bytes
	tagFields := []byte{0}             // 1 Byte
*/

const (
	numbersOfApiKeys = 3
)

var (
	tagFields    = [1]byte{0}
	throttleTime = [4]byte{0, 0, 0, 0}
)

type BodyResp struct {
	ApiKey     int16
	MinVersion int16
	MaxVersion int16
}

var predefinedAPIKeys = []BodyResp{
	{
		// API_Versions
		ApiKey:     18,
		MinVersion: 0,
		MaxVersion: 4,
	},
	{
		// Fetch API
		ApiKey:     1,
		MinVersion: 0,
		MaxVersion: 16,
	},
}

func APIVersions(key uint16) []byte {

	resp := []byte{numbersOfApiKeys}

	for _, apiKey := range predefinedAPIKeys {
		resp = append(resp, AppendBodyResp(apiKey)...)
	}

	// resp = append(resp, tagFields[:]...)
	resp = append(resp, throttleTime[:]...)
	resp = append(resp, tagFields[:]...)

	return resp
}

func AppendBodyResp(resp BodyResp) []byte {
	buffer := new(bytes.Buffer)

	binary.Write(buffer, binary.BigEndian, resp.ApiKey)
	binary.Write(buffer, binary.BigEndian, resp.MinVersion)
	binary.Write(buffer, binary.BigEndian, resp.MaxVersion)

	binary.Write(buffer, binary.BigEndian, tagFields)

	return buffer.Bytes()
}
