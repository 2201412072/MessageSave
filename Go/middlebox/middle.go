package middlebox

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	"mydatabase"
	"os"

	"gorm.io/gorm"
)

var MyDB *gorm.DB
var public_key_path string = "data/master-public.pem"
var private_key_path string = "data/master-private.pem"
var database_path string = "database/mysql"
var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func change_public_key_path(key_path string) int {
	public_key_path = key_path
	return 1
}

func change_private_key_path(key_path string) int {
	private_key_path = key_path
	return 1
}

func change_database_path(db_path string) int {
	database_path = db_path
	return 1
}

func print_all_path() int {
	fmt.Printf("公钥地址:%s\n", public_key_path)
	fmt.Printf("私钥地址:%s\n", private_key_path)
	fmt.Printf("数据库地址:%s\n", database_path)
	return 1
}

func create_myself_key() int {
	temp_privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return 0
	}
	// 从私钥中导出公钥
	temp_publicKey := &temp_privateKey.PublicKey

	// // 原始数据
	// plaintext := []byte("Hello, World!")

	// // 使用公钥加密数据
	// ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, temp_publicKey, plaintext)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("加密后的数据：%x\n", ciphertext)

	// // 使用私钥解密数据
	// decryptedText, err := rsa.DecryptPKCS1v15(rand.Reader, temp_privateKey, ciphertext)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("解密后的数据：%s\n", decryptedText)

	// 将私钥和公钥保存到文件
	privateKeyFile, err := os.Create(private_key_path)
	if err != nil {
		panic(err)
	}
	defer privateKeyFile.Close()
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(temp_privateKey),
	}
	err = pem.Encode(privateKeyFile, privateKeyPEM)
	if err != nil {
		return 0
	}

	publicKeyFile, err := os.Create(public_key_path)
	if err != nil {
		panic(err)
	}
	defer publicKeyFile.Close()
	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(temp_publicKey),
	}
	err = pem.Encode(publicKeyFile, publicKeyPEM)
	if err != nil {
		return 0
	}
	return 1
}

func init_procedure() int {
	//判断是否有公钥、私钥和数据库文件，如果没有就创建一个；否则就加载进去，用于程序一开始的加载
	fmt.Print("程序初始化")
	flag := 1
	privateKeyFile, err := os.Open(private_key_path)
	if err != nil {
		flag = 0
		fmt.Println("私钥地址不存在，准备创建新的公钥和私钥")
	}
	if flag == 1 {
		defer privateKeyFile.Close()

		privateKeyBytes, err := io.ReadAll(privateKeyFile)
		if err != nil {
			//panic(err)
			return 0
		}

		// 解码和解析私钥
		privateKeyBlock, _ := pem.Decode(privateKeyBytes)
		privateKey, err = x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
		if err != nil {
			//panic(err)
			return 0
		}

		fmt.Println("读取并解析私钥成功")
	}

	// 读取公钥文件
	flag2 := 1
	publicKeyFile, err := os.Open(public_key_path)
	if err != nil {
		flag2 = 0
		fmt.Println("私钥地址不存在，准备创建新的公钥和私钥")
	}
	if flag2 == 1 {
		defer publicKeyFile.Close()

		publicKeyBytes, err := io.ReadAll(publicKeyFile)
		if err != nil {
			panic(err)
		}

		// 解码和解析公钥
		publicKeyBlock, _ := pem.Decode(publicKeyBytes)
		publicKey, err = x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
		if err != nil {
			panic(err)
		}

		fmt.Println("读取并解析公钥成功:")
	}
	if flag == 0 || flag2 == 0 {
		flag = create_myself_key()
		if flag != 1 {
			return 0
		}
	}
	MyDB, flag = mydatabase.Db_load(database_path)

	return 1

}

func load() int {
	//判断是否有公钥、私钥和数据库文件，如果有就加载，否则返回错误码
	/*
		err 0：一些系统错误
			1：正常加载
			2：私钥地址不存在
			3：公钥地址不存在
			4：数据库地址不存在
	*/
	flag := 1
	privateKeyFile, err := os.Open(private_key_path)
	if err != nil {
		return 2
	}
	defer privateKeyFile.Close()

	privateKeyBytes, err := io.ReadAll(privateKeyFile)
	if err != nil {
		//panic(err)
		return 0
	}

	// 解码和解析私钥
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	privateKey, err = x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		//panic(err)
		return 0
	}

	fmt.Println("读取并解析私钥成功:")

	// 读取公钥文件
	flag = 1
	publicKeyFile, err := os.Open(public_key_path)
	if err != nil {
		return 3
	}
	defer publicKeyFile.Close()

	publicKeyBytes, err := io.ReadAll(publicKeyFile)
	if err != nil {
		return 0
	}

	// 解码和解析公钥
	publicKeyBlock, _ := pem.Decode(publicKeyBytes)
	publicKey, err = x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return 0
	}

	fmt.Println("读取并解析公钥成功:")
	_, err = os.Stat(database_path)
	if os.IsNotExist(err) {
		return 4
	}
	MyDB, flag = mydatabase.Db_load(database_path)

	return flag

}

func utf_string2base_bytes(message string) ([]byte, int) {
	message_byte := []byte(message)
	base_bytes := base64.StdEncoding.EncodeToString(message_byte)
	return []byte(base_bytes), 1
}

func bytes2base_string(bytes_data []byte) (string, int) {
	encoded := base64.StdEncoding.EncodeToString(bytes_data)
	return encoded, 1
}

func base_string2bytes(message string) ([]byte, int) {
	decoded, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		fmt.Println("Decoding error:", err)
		return make([]byte, 0), 0
	}
	return decoded, 1
}

func base_bytes2utf_string(bytes_data []byte) (string, int) {
	message := string(bytes_data)
	utf_bytes_data, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		fmt.Println("Decoding error:", err)
		return "", 0
	}
	return string(utf_bytes_data), 1
}

func block_encrypt(bytes_data []byte, public_key *rsa.PublicKey) ([]byte, int) {

}
