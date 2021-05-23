package utils

import (
	"bytes"
	"encoding/json"
)

func JsonToByte(t interface{}) []byte {
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(t)
	if err != nil {
		return []byte{}
	}
	return buffer.Bytes()
}

func JsonToString(t interface{}) string {
	buffer := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffer)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(t)
	if err != nil {
		return ""
	}
	return buffer.String()
}
