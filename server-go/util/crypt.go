package util

import (
	"crypto/rand"
	"crypto/rsa"
)

var chunkSize = 256 - 11

// 加密UTF字符串
func EncryptUTFString(message string, public_key *rsa.PublicKey) ([]byte, int) {
	message_byte, _ := UTF_string2base_bytes(message)
	message_encrypted, _ := Block_encrypt(message_byte, public_key)
	return message_encrypted, 1
}

// 加密Base64字符串
func EncryptBase64String(message string, public_key *rsa.PublicKey) ([]byte, int) {
	message_byte, _ := Base_string2bytes(message)
	message_encrypted, _ := Block_encrypt(message_byte, public_key)
	return message_encrypted, 1
}

// 使用公钥对明文加密（已经考虑了分块加密）
func Block_encrypt(bytes_data []byte, public_key *rsa.PublicKey) ([]byte, int) {
	ciphertext := make([]byte, 0)

	for i := 0; i < len(bytes_data); i += chunkSize {
		endIndex := i + chunkSize
		if endIndex > len(bytes_data) {
			endIndex = len(bytes_data)
		}

		chunk, err := rsa.EncryptPKCS1v15(rand.Reader, public_key, bytes_data[i:endIndex])
		if err != nil {
			return nil, 0
		}

		ciphertext = append(ciphertext, chunk...)
	}
	return ciphertext, 1
}

// 使用私钥对密文解密（同理，这个为分块解密）
func Block_decrypt(bytes_data []byte, private_key *rsa.PrivateKey) ([]byte, int) {
	//
	plaintext := make([]byte, 0)

	for i := 0; i < len(bytes_data); i += private_key.Size() {
		endIndex := i + private_key.Size()
		if endIndex > len(bytes_data) {
			endIndex = len(bytes_data)
		}

		chunk, err := rsa.DecryptPKCS1v15(rand.Reader, private_key, bytes_data[i:endIndex])
		if err != nil {
			return nil, 0
		}

		plaintext = append(plaintext, chunk...)
	}

	return plaintext, 1
}
