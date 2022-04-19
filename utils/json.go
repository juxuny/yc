package utils

import "encoding/json"

func ToJson(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}
