package tool

import (
	"bytes"
	"encoding/json"
)

/**
 * @PROJECT_NAME wechat
 * @author  Moqi
 * @date  2021-04-23 10:37
 * @Email:str@li.cm
 **/

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
