package controller

import (
	"client-go/model"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"os"
)

var public_key_path string = "./data/master-public.pem"
var private_key_path string = "./data/master-private.pem"
var Private_key *rsa.PrivateKey
var Public_key *rsa.PublicKey
var username string = "test1"
var server_addr string = "127.0.0.1"
var port string = "8091"
var ResponseChan = make(chan string) //该通道用于子协程传递回应消息

// 修改公钥保存路径
func Change_public_key_path(key_path string) int {
	public_key_path = key_path
	return 1
}

// 修改私钥保存路径
func Change_private_key_path(key_path string) int {
	private_key_path = key_path
	return 1
}

// // 修改数据库存储路径
// func Change_database_path(db_path string) int {
// 	database_path = db_path
// 	return 1
// }

// 打印存储数据地址
func Print_all_path() int {
	fmt.Printf("公钥地址:%s\n", public_key_path)
	fmt.Printf("私钥地址:%s\n", private_key_path)
	// fmt.Printf("数据库地址:%s\n", database_path)
	return 1
}

// 创建用户私钥与公钥
func Create_myself_key() int {
	temp_privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return 0
	}
	// 从私钥中导出公钥
	temp_publicKey := &temp_privateKey.PublicKey

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

// 加载或新建公钥私钥
func LoadKeys() int {
	// 创建数据文件夹
	_, err := os.Stat("./data")
	if os.IsNotExist(err) {
		// 文件夹不存在，使用 os.MkdirAll 函数创建文件夹
		err := os.MkdirAll("./data", os.ModePerm)
		if err != nil {
			fmt.Println("创建文件夹失败:", err)
			return 0
		}
	}

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
		Private_key, err = x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
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
		Public_key, err = x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
		if err != nil {
			panic(err)
		}

		fmt.Println("读取并解析公钥成功:")
	}
	if flag == 0 || flag2 == 0 {
		flag = Create_myself_key()
		if flag != 1 {
			return 0
		}
	}
	return 1
}

func PublicKeyBytesFromFile() ([]byte, int) {
	//该函数用于取出本地公钥的byte形式，用于向服务器传递
	publicKeyFile, err := os.Open(public_key_path)
	if err != nil {
		//公钥地址不存在
		return make([]byte, 0), 0
	}
	defer publicKeyFile.Close()

	publicKeyBytes, err := io.ReadAll(publicKeyFile)
	if err != nil {
		//读取公钥文件出现错误
		return make([]byte, 0), 2
	}
	return publicKeyBytes, 1

}

func Init() {
	// 加载配置文件
	// 加载公钥私钥
	LoadKeys()
	// 加载数据库？
	model.SetupDB()
}
