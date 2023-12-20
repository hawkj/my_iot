package common

import "encoding/json"

func IsValidJSON(s string) bool {
	// 将字符串转换成字节切片
	data := []byte(s)

	// 使用json.Valid判断是否是有效的JSON
	return json.Valid(data)
}
