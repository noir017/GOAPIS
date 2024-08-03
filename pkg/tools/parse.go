package tools

import (
	"encoding/json"
	"fmt"
)

func StrToJson(s string) (map[string]interface{}, error) {

	var data map[string]interface{}
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		fmt.Println("StrToJson Error:", err)
		return nil, err
	}
	return data, nil
}
