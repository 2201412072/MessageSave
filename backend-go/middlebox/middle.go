package middlebox

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io"
	mrand "math/rand"
	"mydatabase"
	"os"
	"strings"
	"time"

	"gorm.io/gorm"
)

var MyDB *gorm.DB
var public_key_path string = "./data/master-public.pem"
var private_key_path string = "./data/master-private.pem"
var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

var layout = "2006-01-02 15:04:05"
var chunkSize = 256 - 11

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

// 加载或新建公钥私钥与数据库
func Init_procedure() int {
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
		flag = Create_myself_key()
		if flag != 1 {
			return 0
		}
	}
	MyDB, flag = mydatabase.Db_load()

	return 1

}

// 加载公钥私钥和数据库
func Load() int {
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
	// _, err = os.Stat(database_path)
	// if os.IsNotExist(err) {
	// 	return 4
	// }
	MyDB, flag = mydatabase.Db_load()

	return flag

}

// 字节类型转公钥
func publicKey_from_bytes(publicKeyBytes []byte) (*rsa.PublicKey, int) {
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

// 使用公钥对明文加密
func Block_encrypt(bytes_data []byte, public_key *rsa.PublicKey) ([]byte, int) {
	//
	ciphertext := make([]byte, 0)

	for i := 0; i < len(bytes_data); i += chunkSize {
		endIndex := i + chunkSize
		if endIndex > len(bytes_data) {
			endIndex = len(bytes_data)
		}

		chunk, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, bytes_data[i:endIndex])
		if err != nil {
			return nil, 0
		}

		ciphertext = append(ciphertext, chunk...)
	}
	return ciphertext, 1
}

// 使用私钥对密文解密
func Block_decrypt(bytes_data []byte) ([]byte, int) {
	//
	plaintext := make([]byte, 0)

	for i := 0; i < len(bytes_data); i += privateKey.Size() {
		endIndex := i + privateKey.Size()
		if endIndex > len(bytes_data) {
			endIndex = len(bytes_data)
		}

		chunk, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, bytes_data[i:endIndex])
		if err != nil {
			return nil, 0
		}

		plaintext = append(plaintext, chunk...)
	}

	return plaintext, 1
}

// 关联用户对数据加密
func Using_other_public(username string, bytes_data []byte) ([]byte, int) {
	temp, flag := mydatabase.Get_single_public_key(username)
	if flag == 0 {
		return nil, 0
	}
	// return temp.Public_key,1
	temp_key, flag := publicKey_from_bytes(temp.Public_key)
	if flag != 1 {
		return nil, 0
	}
	cipher_key, flag := Block_encrypt(bytes_data, temp_key)
	if flag != 1 {
		return nil, 0
	}
	return cipher_key, 1
}

// 对数据加密
func Using_myself_public(bytes_data []byte) ([]byte, int) {
	cipher_key, flag := Block_encrypt(bytes_data, publicKey)
	if flag != 1 {
		return nil, 0
	}
	return cipher_key, 1
}

// 对密文解密
func Using_myself_private(bytes_data []byte) ([]byte, int) {
	cipher_key, flag := Block_decrypt(bytes_data)
	if flag != 1 {
		return nil, 0
	}
	return cipher_key, 1
}

// 生成操作请求
func Agree_A2B_stage1(operator string) (string, int) {
	nowtime := time.Now()
	timeStr := nowtime.Format(layout)
	agree_message := timeStr + "_" + operator
	return agree_message, 1
}

// 验证操作请求是否同意
func Agree_A2B_stage2(username string, mytime time.Time, message string) int {
	temp_bytes_data, flag := Base_string2bytes(message)
	if flag == 0 {
		return 0
	}
	temp_public_key_bytes, flag := mydatabase.Get_single_public_key(username)
	if flag == 0 {
		return 0
	}
	temp_public_key, flag := publicKey_from_bytes(temp_public_key_bytes.Public_key)
	if flag == 0 {
		return 0
	}
	temp_op, flag := mydatabase.Get_single_user_time_agree(username, mytime)
	if flag != 1 {
		return 0
	}
	timeStr := mytime.Format(layout)
	verify_string := "agree_" + timeStr + "_" + temp_op.Operator
	verify_bytes, flag := UTF_string2base_bytes(verify_string)
	if flag != 1 {
		return 0
	}
	hash := sha256.Sum256(verify_bytes)
	err := rsa.VerifyPKCS1v15(temp_public_key, crypto.SHA256, hash[:], temp_bytes_data)
	if err != nil {
		fmt.Println("Signature verification failed:", err)
		return 0
	}
	return 1
}

func Agree_B2A_stage1(username string, message string) (string, time.Time, string, int) {
	result := strings.Split(message, "_")
	if len(result) != 2 {
		return "", time.Time{}, "", 0
	}
	temp_time, err := time.Parse(layout, result[0])
	if err != nil {
		return "", time.Time{}, "", 0
	}
	return username, temp_time, result[1], 1
}

func agree_B2A_stage2_a(mytime time.Time, operator string) (string, int) {
	timeStr := mytime.Format(layout)
	agree_string := "agree_" + timeStr + "_" + operator
	temp_bytes_data, flag := UTF_string2base_bytes(agree_string)
	if flag != 1 {
		return "", 0
	}
	hash := sha256.Sum256(temp_bytes_data)
	bytes_data, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", 0
	}
	string_data, flag := Bytes2base_string(bytes_data)
	if flag != 1 {
		return "", 0
	}
	return string_data, 1
}

func agree_B2A_stage2_d() (string, int) {
	agree_string := "don't agree"
	temp_bytes_data, flag := UTF_string2base_bytes(agree_string)
	if flag != 1 {
		return "", 0
	}
	hash := sha256.Sum256(temp_bytes_data)
	bytes_data, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return "", 0
	}
	string_data, flag := Bytes2base_string(bytes_data)
	if flag != 1 {
		return "", 0
	}
	return string_data, 1
}

// 对Base64 string类型的数据进行一次解密得到UTF string类型的数据（整个解密过程需进行两次该操作，一个用户一次）
func Deal_B2A_message(message string) (string, int) {
	temp_bytes_data, flag := Base_string2bytes(message)
	if flag != 1 {
		return "", 1
	}
	bytes_data, _ := Using_myself_private(temp_bytes_data)
	string_data, flag := Base_bytes2utf_string(bytes_data)
	if flag != 1 {
		return "", 0
	}
	return string_data, 1
}

// 读取公钥
func Show_myself_public_key() (string, int) {
	// 读取公钥文件
	publicKeyFile, err := os.Open(public_key_path)
	if err != nil {
		return "", 0
	}
	defer publicKeyFile.Close()

	publicKeyBytes, err := io.ReadAll(publicKeyFile)
	if err != nil {
		return "", 0
	}
	context, flag := Bytes2base_string(publicKeyBytes)
	if flag != 1 {
		return "", 0
	}
	return context, 1
}

// 读取所有用户公钥
func Get_all_public_key() ([]string, []string, int) {
	result, flag := mydatabase.Get_all_public_keys()
	if flag != 1 {
		return make([]string, 0), make([]string, 0), 0
	}
	var arr_username []string
	var arr_public_key []string
	for _, item := range result {
		string_data, flag := Bytes2base_string(item.Public_key)
		if flag != 1 {
			continue
		}
		arr_username = append(arr_username, item.Username)
		arr_public_key = append(arr_public_key, string_data)
	}
	return arr_username, arr_public_key, 1
}

// 读取指定用户的公钥
func Get_single_public_key(username string) (string, int) {
	result, flag := mydatabase.Get_single_public_key(username)
	if flag != 1 {
		return "", 0
	}
	string_data, flag := Bytes2base_string(result.Public_key)
	if flag != 1 {
		return "", 0
	}
	return string_data, 1
}

// 加入其他用户公钥
func Add_other_public_key(othername string, other_public_key string) int {
	bytes_data, flag := Base_string2bytes(other_public_key)
	if flag != 1 {
		return 0
	}
	mydatabase.Add_public_key(othername, bytes_data)
	return 1
}

// 发送删除他人公钥请求
func Delete_other_public_key_stage1(othername string) (string, int) {
	return Agree_A2B_stage1("delete")
}

// 对方同意后删除其公钥
func Delete_other_public_key_stage2(othername string, mytime time.Time, message string) int {
	result := Agree_A2B_stage2(othername, mytime, message)
	if result == 1 {
		mydatabase.Delete_public_key(othername)
		return 1
	}
	return 0
}

// 更改他人公钥
func Change_other_public_key(username string, public_key_message string) int {
	public_key, flag := Base_string2bytes(public_key_message)
	if flag != 1 {
		return 0
	}
	mydatabase.Change_public_key(username, public_key)
	return 1
}

// 对数据进行双重加密，并随机选择一个提示字符
func password2m_s_random(username string, password string) ([]byte, string, int) {
	temp, _ := UTF_string2base_bytes(password)
	multiply_key, _ := Using_myself_public(temp)
	bytes_data, flag := Using_other_public(username, multiply_key)
	if flag != 1 {
		return make([]byte, 0), "", 0
	}
	randomIndex := mrand.Intn(len(password))
	single_key := string(password[randomIndex])
	return bytes_data, single_key, 1
}

// 对数据进行双重加密，并指定提示字符
func password2m_s_choice(username string, password string, single_char string) ([]byte, string, int) {
	//2表示single char长度不为1，或者不在password里
	temp, _ := UTF_string2base_bytes(password)
	multiply_key, _ := Using_myself_public(temp)
	bytes_data, flag := Using_other_public(username, multiply_key)
	if flag != 1 {
		return make([]byte, 0), "", 0
	}
	if len(single_char) == 1 && strings.Contains(password, single_char) {
		return bytes_data, single_char, 1
	} else {
		return bytes_data, "", 2
	}
}

// 添加密码
func Add_password(application string, username string, password string, mode string, single_char string) int {
	var saved_key []byte
	var single_key string
	flag := 0
	if mode == "random" {
		saved_key, single_key, flag = password2m_s_random(username, password)
	} else if mode == "single" {
		saved_key, single_key, flag = password2m_s_choice(username, password, single_char)
	} else {
		return 0
	}
	if flag != 1 {
		return 0
	}
	mydatabase.Add_password(application, username, saved_key, single_key)
	return 1
}

// 删除密码
func Delete_single_password(application string, username string) int {
	return mydatabase.Delete_single_password(application, username)
}

// 删除与指定用户关联的所有密码
func Delete_username_password(username string) int {
	return mydatabase.Delete_username_password(username)
}

// 获得指定应用对应的关联用户
func Get_application2user_password(application string) ([]string, int) {
	result, flag := mydatabase.Get_application2user_password(application)
	if flag != 1 {
		return make([]string, 0), 0
	}
	var arr_username []string
	for _, item := range result {
		arr_username = append(arr_username, item.Username)
	}
	return arr_username, 1
}

// 获得指定应用与指定关联用户的密码与提示字符
func Get_application_name2key_password(applicatoin string, username string) (string, string, int) {
	result, flag := mydatabase.Get_application_name2key_password(applicatoin, username)
	if flag != 1 {
		return "", "", 0
	}
	string_data, flag := Bytes2base_string(result.Saved_key)
	if flag != 1 {
		return "", result.Single_key, 2
	} else {
		return string_data, result.Single_key, 1
	}
}

// 添加重要信息
func Add_important(keyword string, username string, important_message string) int {
	saved_key, _, flag := password2m_s_random(username, important_message)
	if flag != 1 {
		return 0
	}
	flag = mydatabase.Add_important(keyword, username, saved_key)
	if flag != 1 {
		return 0
	}
	return 1
}

// 获得所有重要信息的关键词
func Get_keyword_important() ([]string, int) {
	result, flag := mydatabase.Get_keyword_important()
	if flag != 1 {
		return make([]string, 0), 0
	}
	var arr_keyword []string
	for _, item := range result {
		arr_keyword = append(arr_keyword, item.Keyword)
	}
	return arr_keyword, 1
}

// 获得重要信息表中指定关键词对应的关联用户
func Get_keyword2user_important(keyword string) ([]string, int) {
	result, flag := mydatabase.Get_keyword2user_important(keyword)
	if flag != 1 {
		return make([]string, 0), 0
	}
	var arr_usernaem []string
	for _, item := range result {
		arr_usernaem = append(arr_usernaem, item.Username)
	}
	return arr_usernaem, 1
}

// 获得重要信息表中指定关联用户对应的所有关键词
func Get_user2keyword_important(username string) ([]string, int) {
	result, flag := mydatabase.Get_user2keyword_important(username)
	if flag != 1 {
		return make([]string, 0), 0
	}
	var arr_keyword []string
	for _, item := range result {
		arr_keyword = append(arr_keyword, item.Keyword)
	}
	return arr_keyword, 1
}

// 获得重要信息表中指定关联用户与指定关键词对应的重要信息
func Get_keyword_name2key_important(keyword string, uername string) (string, int) {
	result, flag := mydatabase.Get_keyword_name2key_important(keyword, uername)
	if flag != 1 {
		return "", 0
	}
	string_data, flag := Bytes2base_string(result.Saved_key)
	if flag != 1 {
		return "", 0
	}
	return string_data, 1
}

// 删除指定关联用户与指定关键词对应的重要信息
func Delete_single_important(keyword string, username string) int {
	return mydatabase.Delete_single_important(keyword, username)
}

// 删除重要信息表中指定关联用户的所有重要信息
func Delete_username_important(username string) int {
	return mydatabase.Delte_username_important(username)
}

// 删除重要信息表中指定关键词对应的所有重要信息
func Delete_keyword_important(keyword string) int {
	return mydatabase.Delete_keyword_important(keyword)
}
