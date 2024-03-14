package controller

import (
	"client-go/util"
)

//该文件对应解密过程

// 对Base64 string类型的数据进行一次解密得到base string类型的数据（整个解密过程需进行一次该操作，B用户收到A请求后）
func Deal_B2A_message_to_base(message string) (string, int) {
	temp_bytes_data, flag := util.Base_string2bytes(message)
	if flag != 1 {
		return "", 1
	}
	bytes_data, _ := util.Block_decrypt(temp_bytes_data, Private_key)
	string_data, flag := util.Bytes2base_string(bytes_data)
	if flag != 1 {
		return "", 0
	}
	return string_data, 1
}

// 对Base64 string类型的数据进行一次解密得到UTF string类型的数据（整个解密过程需进行一次该操作，A用户收到B回复后）
func Deal_B2A_message_to_utf(message string) (string, int) {
	temp_bytes_data, flag := util.Base_string2bytes(message)
	if flag != 1 {
		return "", 1
	}
	bytes_data, _ := util.Block_decrypt(temp_bytes_data, Private_key)
	string_data, flag := util.Base_bytes2utf_string(bytes_data)
	if flag != 1 {
		return "", 0
	}
	return string_data, 1
}
