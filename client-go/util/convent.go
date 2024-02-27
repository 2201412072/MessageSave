package util

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

// 字节类型转公钥
func PublicKey_from_bytes(publicKeyBytes []byte) (*rsa.PublicKey, int) {
	publicKeyBlock, _ := pem.Decode(publicKeyBytes)
	//var temp *rsa.PublicKey
	temp, err := x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}
	return temp, 1
}

// UTF字符串转Base64字节
func UTF_string2base_bytes(message string) ([]byte, int) {
	message_byte := []byte(message)
	base_bytes := base64.StdEncoding.EncodeToString(message_byte)
	return []byte(base_bytes), 1
}

// Base64字节转Base字符串
func Bytes2base_string(bytes_data []byte) (string, int) {
	encoded := base64.StdEncoding.EncodeToString(bytes_data)
	return encoded, 1
}

// Base字符串转Base64字节
func Base_string2bytes(message string) ([]byte, int) {
	decoded, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		fmt.Println("Decoding error:", err)
		return make([]byte, 0), 0
	}
	return decoded, 1
}

// Base64字节转UTF字符串
func Base_bytes2utf_string(bytes_data []byte) (string, int) {
	message := string(bytes_data)
	// message := base64.StdEncoding.EncodeToString(bytes_data)
	utf_bytes_data, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		fmt.Println("Decoding error:", err)
		return "", 0
	}
	return string(utf_bytes_data), 1
}
