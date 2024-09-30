package protocol

import (
	"bytes"
	"encoding/binary"
)

/*
  - .ResponseBody
	✔️ .error_code (0)
  ✔️ .num_api_keys (2)
  - .ApiKeys[0]
    ✔️ .api_key (18)
    ✔️ .min_version (0)
    ✔️ .max_version (4)
    ✔️ .TAG_BUFFER
  - .ApiKeys[1]
    ✔️ .api_key (1)
    ✔️ .min_version (0)
    ✔️ .max_version (16)
    ✔️ .TAG_BUFFER
    ✔️ .throttle_time_ms (0)
    ✔️ .TAG_BUFFER
	✔️ .throttle_time_ms (0)
  ✔️ .TAG_BUFFER
*/

type APIVersionsBody struct {
	ApiKey     int16
	MinVersion int16
	MaxVersion int16
	TagBuffer  int8
}

var predefinedAPIKeys = []APIVersionsBody{
	{
		// API_Versions
		ApiKey:     18,
		MinVersion: 0,
		MaxVersion: 4,
		TagBuffer:  0,
	},
	{
		// Fetch API
		ApiKey:     1,
		MinVersion: 0,
		MaxVersion: 16,
		TagBuffer:  0,
	},
}

func ApiVersionBodyResp(apiKey, apiVersion uint16) []byte {
	bodyBuffer := new(bytes.Buffer)

	binary.Write(bodyBuffer, binary.BigEndian, ErrorCodeChecker(apiKey, apiVersion))
	binary.Write(bodyBuffer, binary.BigEndian, NUMS_API_KEYS)

	for _, apiKey := range predefinedAPIKeys {
		apiVersionsBodyChecker(bodyBuffer, apiKey)
	}

	binary.Write(bodyBuffer, binary.BigEndian, THROTTLE_TIME)
	binary.Write(bodyBuffer, binary.BigEndian, TAG_BUFFER)

	return bodyBuffer.Bytes()
}

func apiVersionsBodyChecker(bodyBuffer *bytes.Buffer, version APIVersionsBody) {
	binary.Write(bodyBuffer, binary.BigEndian, version.ApiKey)
	binary.Write(bodyBuffer, binary.BigEndian, version.MinVersion)
	binary.Write(bodyBuffer, binary.BigEndian, version.MaxVersion)
	binary.Write(bodyBuffer, binary.BigEndian, version.TagBuffer)
}
