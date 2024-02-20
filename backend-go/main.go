package main

import (
	"crypto/rsa"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"middlebox"
	"mydatabase"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey
var db *gorm.DB

type PostParams struct {
	NKey_word string `json:"key_word"`
	Passwd    string `json:"passwd"`
	User      string `json:"user"`
}

func main() {
	// 加载公钥私钥
	middlebox.Init_procedure()
	// middlebox.Load()
	// 加载数据库
	// mydatabase.Db_load()
	// db, err := mydatabase.Db_load()
	// if err != 1 {
	// 	return
	// }

	r := gin.Default()
	r.Use(middlebox.CORSHandler()) // 设置全局跨域访问
	r.POST("/testBind", func(c *gin.Context) {
		var p PostParams
		err := c.ShouldBindJSON(&p) //将传过来的json赋给p
		if err != nil {
			fmt.Println(err)
			c.JSON(200, gin.H{
				"msg":  "bindjson报错了",
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "请求成功",
				"data": p, //如果请求成功则把p的值打印出来
			})
		}
	})

	// HomePage
	r.POST("/SearchPassword", func(ctx *gin.Context) { // 查询密码
		key_word := ctx.PostForm("key_word")
		user := ctx.PostForm("connect_user")
		// 检索数据库
		passwd, _, err := middlebox.Get_application_name2key_password(key_word, user)
		if err != 1 {
			return
		}
		// 将密码发送给用户B进行一次解密
		// 对解密后密码进行二次解密
		decrypted_passwd, err := middlebox.Deal_B2A_message(passwd)
		// 返回结果
		ctx.JSON(200, gin.H{"passwd": decrypted_passwd})
	})

	r.POST("/AddPassword", func(ctx *gin.Context) { // 添加密码
		// 前端传的是json，所以不能简单地用PostForm解析。可以用如下代码
		// var requestMap1 = make(map[string]string)
		// json.NewDecoder(ctx.Request.Body).Decode(&requestMap1)
		/*
			key_word := ctx.PostForm("key_word")
			passwd := ctx.PostForm("passwd")
			user := ctx.PostForm("user")
		*/

		var requestMap struct {
			Key_word string `json:"key_word"`
			Passwd   string `json:"passwd"`
			User     string `json:"user"`
		}
		err := ctx.ShouldBind(&requestMap)
		if err != nil {
			fmt.Println("解析请求体失败:", err)
			ctx.String(http.StatusNotFound, "绑定form失败")
		}
		fmt.Println("AddPassword", requestMap.Key_word, requestMap.Passwd, requestMap.User)
		// 判断用户是否存在
		// 将密码加密，需双人加密才行!
		flag := middlebox.Add_password(requestMap.Key_word, requestMap.User, requestMap.Passwd, "random", "")

		// fmt.Println("AddPassword", requestMap.Key_word, requestMap.Passwd, requestMap.User)
		// // 判断用户是否存在
		// // 将密码加密，需双人加密才行!
		// flag := middlebox.Add_password(requestMap.Key_word, requestMap.User, requestMap.Passwd, "random", "")
		if flag != 1 {
			return
		}
		// 返回结果
		ctx.JSON(200, gin.H{"msg": "successfully saved."})
	})

	// PassswordManage
	r.GET("/PasswordManage", func(ctx *gin.Context) {
		// 获取现有保存的密码数据
		passwd_data, err := mydatabase.Get_password()
		if err != 1 {
			return
		}
		// passwd_data_json,err := json.Marshal(passwd_data)
		// if err != nil{
		// 	ctx.JSON(200, gin.H{"msg": "Null"})
		// 	return
		// }
		ctx.JSON(200, passwd_data)
	})

	// MessageManage
	// r.GET()

	// r.Run()
	r.Run("127.0.0.1:8090")
}
