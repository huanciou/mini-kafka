package protocol

import (
	"bytes"
	"encoding/binary"
)

/*
tester check:

- .ResponseHeader
		✔️ .correlation_id (1314877239)
		✔️ .TAG_BUFFER

- .ResponseBody
		✔️ .throttle_time_ms (587202560)
		✔️ .error_code (0)
		✔️ .session_id (0)
		✔️ .num_responses (0)
		✔️ .TAG_BUFFER
*/

func FetchAPIBodyResp(apiKey, apiVersion uint16) []byte {
	bodyBuffer := new(bytes.Buffer)

	binary.Write(bodyBuffer, binary.BigEndian, THROTTLE_TIME)

	binary.Write(bodyBuffer, binary.BigEndian, ErrorCodeChecker(apiKey, apiVersion))

	binary.Write(bodyBuffer, binary.BigEndian, SESSION_ID)
	binary.Write(bodyBuffer, binary.BigEndian, NUMS_RESPONSES)
	binary.Write(bodyBuffer, binary.BigEndian, TAG_BUFFER)

	return bodyBuffer.Bytes()
}
