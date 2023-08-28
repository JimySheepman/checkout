package utils

import "encoding/json"

func ToJSON(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		panic("json: " + err.Error())
	}

	return string(data)
}
