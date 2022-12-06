package utils

import "encoding/json"

// AnyToString 将任意类型转换成String
func AnyToString(v interface{}) string {
	value, _ := json.Marshal(v)
	return string(value)
}

// AnyToStringBeauty 将任意类型转换成带缩进格式的String
func AnyToStringBeauty(v interface{}) string {
	value, _ := json.MarshalIndent(v, "", "    ")
	return string(value)
}
