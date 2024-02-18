package main

import (
	"crypto/rsa"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"middlebox"
	"mydatabase"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey
var db *gorm.DB
var database_path string = "database/mysql"

func main() {
	// 加载公钥私钥
	middlebox.Load()
	// 加载数据库
	mydatabase.Db_load(database_path)
	// db, err := mydatabase.Db_load()
	// if err != 1 {
	// 	return
	// }

	r := gin.Default()
	// HomePage
	r.POST("/SearchPassword", func(ctx *gin.Context) { // 查询密码
		key_word := ctx.PostForm("key_word")
		user := ctx.PostForm("connect_user")
		// 检索数据库
		password_struct, err := mydatabase.Get_application_name2key_password(key_word, user)
		if err != 1 {
			return
		}
		passwd := password_struct.Saved_key
		// 将密码解密
		// passwd_bytes, err := middlebox.Base_string2bytes(passwd)
		decrypted_passwd, err := middlebox.Block_decrypt(passwd, privateKey)
		// 返回结果
		ctx.JSON(200, gin.H{"passwd": decrypted_passwd})
	})

	r.POST("/AddPassword", func(ctx *gin.Context) { // 添加密码
		key_word := ctx.PostForm("key_word")
		passwd := ctx.PostForm("passwd")
		user := ctx.PostForm("user")
		fmt.Println("AddPassword", key_word, passwd, user)
		// 判断用户是否存在
		// 将密码加密，需双人加密才行!
		passwd_bytes, err := middlebox.Base_string2bytes(passwd)
		if err != 1 {
			return
		}
		encrypted_passwd, err := middlebox.Block_encrypt(passwd_bytes, publicKey)
		// encrypted_passwd_bytes,err := middlebox.Base_bytes2utf_string(encrypted_passwd)
		// 添加该密码
		err = mydatabase.Add_password(key_word, user, encrypted_passwd, "")
		if err != 1 {
			return
		}
		// 返回结果
		ctx.JSON(200, gin.H{"msg": "successfully saved."})
	})

	// PassswordManage
	r.GET("/PasswordManage", func(ctx *gin.Context) {
		// 获取现有保存的密码数据
		ctx.JSON(200, gin.H{"msg": "Null"})
	})

	// MessageManage
	// r.GET()

	// r.Run()
	r.Run(":8090")
}
