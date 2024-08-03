package extapi

import (
	"errors"
)

// 获取每日新闻接口，格式化接口内容

// https://api.vvhan.com/api转换
func ParseVvhan(data map[string]interface{}) ([]Hotword, error) {
	var ret []Hotword
	if apidatas, ok1 := (data)["data"].([]interface{}); ok1 {
		for _, v := range apidatas {
			// datemap := v.(map[string]interface{})
			if item, ok := v.(map[string]interface{}); ok {
				temp := Hotword{
					Title:  item["title"].(string),
					Detail: item["url"].(string),
				}
				ret = append(ret, temp)
			} else {
				return ret, errors.New("ParseVvhan error")
			}
		}
	} else {
		return ret, errors.New("ParseVvhan error")
	}

	return ret, nil
}
