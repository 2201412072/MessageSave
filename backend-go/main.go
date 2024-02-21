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
	temp, _ := middlebox.Show_myself_public_key()
	fmt.Println("public key:", temp)

	r := gin.Default()
	r.Use(middlebox.CORSHandler()) // 设置全局跨域访问
	// HomePage
	r.POST("/SearchPassword", func(ctx *gin.Context) { // 查询密码
		var requestMap struct {
			Key_word string `json:"key_word"`
			User     string `json:"user"`
		}
		err := ctx.ShouldBind(&requestMap)
		if err != nil {
			fmt.Println("解析请求体失败:", err)
			ctx.String(http.StatusNotFound, "绑定form失败")
		}
		// 检索数据库
		passwd, _, flag := middlebox.Get_application_name2key_password(requestMap.Key_word, requestMap.User)
		if flag != 1 {
			return
		}
		// 将密码发送给用户B进行一次解密
		// 对解密后密码进行二次解密
		//decrypted_passwd, _ := middlebox.Deal_B2A_message_to_base(passwd)
		// 返回结果
		ctx.JSON(200, gin.H{"passwd": passwd})
	})
	r.POST("/DecryptPassword", func(ctx *gin.Context) { //A收到B的回应消息，得到最终密码
		var requestMap struct {
			Encrypted_password string `json:"encrypted_password"`
		}
		err := ctx.ShouldBind(&requestMap)
		if err != nil {
			fmt.Println("解析请求体失败:", err)
			ctx.String(http.StatusNotFound, "绑定form失败")
		}
		passwd, flag := middlebox.Deal_B2A_message_to_utf(requestMap.Encrypted_password)
		if flag != 1 {
			return
		}
		ctx.JSON(200, gin.H{"passwd": passwd})
	})
	r.POST("/DecryptOtherPassword", func(ctx *gin.Context) { //B收到A发送的解密消息，进行解密
		var requestMap struct {
			Encrypt_message string `json:"encrypted_password"`
		}
		err := ctx.ShouldBind(&requestMap)
		if err != nil {
			fmt.Println("解析请求体失败:", err)
			ctx.String(http.StatusNotFound, "绑定form失败")
		}
		passwd, flag := middlebox.Deal_B2A_message_to_base(requestMap.Encrypt_message)
		if flag != 1 {
			return
		}
		ctx.JSON(200, gin.H{"passwd": passwd})
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
			Key_word string `json:"Key_word"`
			Passwd   string `json:"Passwd"`
			User     string `json:"User"`
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
	r.POST("/AddMessage", func(ctx *gin.Context) { // 添加重要信息
		var requestMap struct {
			Key_word string `json:"key_word"`
			Message  string `json:"message"`
			User     string `json:"user"`
		}
		err := ctx.ShouldBind(&requestMap)
		if err != nil {
			fmt.Println("解析请求体失败:", err)
			ctx.String(http.StatusNotFound, "绑定form失败")
		}
		fmt.Println("AddMessage", requestMap.Key_word, requestMap.Message, requestMap.User)
		// 判断用户是否存在
		// 将密码加密，需双人加密才行!
		flag := middlebox.Add_important(requestMap.Key_word, requestMap.User, requestMap.Message)

		if flag != 1 {
			return
		}
		// 返回结果
		ctx.JSON(200, gin.H{"msg": "successfully saved."})
	})
	r.POST("/AddUserPublicKey", func(ctx *gin.Context) {
		var requestMap struct {
			Username   string `json:"Username"`
			Public_key string `json:"Public_key"`
		}
		err := ctx.ShouldBindJSON(&requestMap)
		if err != nil {
			fmt.Println("解析请求体失败:", err)
			ctx.String(http.StatusNotFound, "绑定form失败")
		}
		contentype := ctx.ContentType()
		fmt.Print("ContentType:", contentype)
		flag := middlebox.Add_other_public_key(requestMap.Username, requestMap.Public_key)
		if flag != 1 {
			ctx.String(http.StatusBadRequest, "")
			return
		}
		// 返回结果
		ctx.JSON(200, gin.H{"msg": "successfully added."})
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
		var result []map[string]interface{}
		for _, item := range passwd_data {
			tempMap := map[string]interface{}{
				"key_word":     item.Application,
				"connect_user": item.Username,
			}
			result = append(result, tempMap)
		}
		ctx.JSON(200, result)
	})
	r.POST("/PasswordManage/Search", func(ctx *gin.Context) {
		//通过关联用户以及应用名字，查询密码
		var requestMap struct {
			Key_word     string `json:"key_word"`
			Connect_user string `json:"connect_user"`
		}
		err := ctx.ShouldBind(&requestMap)
		if err != nil {
			fmt.Println("解析请求体失败:", err)
			ctx.String(http.StatusNotFound, "绑定form失败")
		}
		if requestMap.Key_word != "" && requestMap.Connect_user != "" {
			//两者都有，因此只展示一个
			_, _, flag := middlebox.Get_application_name2key_password(requestMap.Key_word, requestMap.Connect_user)
			if flag != 1 {
				ctx.String(http.StatusBadRequest, "该关联用户未存储该应用密码")
				return
			}
		} else if requestMap.Key_word == "" {
			ctx.String(http.StatusBadRequest, "必须输入应用名")
			return
		} else if requestMap.Key_word != "" {
			//var users []string
			users, flag := middlebox.Get_application2user_password(requestMap.Key_word)
			if flag != 1 || len(users) == 0 {
				ctx.String(http.StatusBadRequest, "无法查到有关联用户存储着该应用密码")
				return
			}
			var result []map[string]interface{}
			for _, item := range users {
				tempMap := map[string]interface{}{
					"key_word":     requestMap.Key_word,
					"connect_user": item,
				}
				result = append(result, tempMap)
			}
			ctx.JSON(200, result)
		}
	})
	r.POST("/PasswordManage/Delete", func(ctx *gin.Context) {
		//通过关联用户以及应用名字，删除该条密码
		var requestMap struct {
			Key_word     string `json:"key_word"`
			Connect_user string `json:"connect_user"`
		}
		err := ctx.ShouldBind(&requestMap)
		if err != nil {
			fmt.Println("解析请求体失败:", err)
			ctx.String(http.StatusNotFound, "绑定form失败")
		}
		if requestMap.Key_word == "" || requestMap.Connect_user == "" {
			ctx.String(http.StatusBadRequest, "信息不全")
			return
		}
		flag := middlebox.Delete_single_password(requestMap.Key_word, requestMap.Connect_user)
		if flag != 1 {
			ctx.String(http.StatusBadRequest, "未删除")
			return
		}
		ctx.JSON(200, gin.H{"msg": "successfully added."})
	})

	// MessageManage
	r.GET("/MessageManage", func(ctx *gin.Context) {
		// 获取现有保存的密码数据
		important_data, err := mydatabase.Get_important()
		if err != 1 {
			return
		}
		// passwd_data_json,err := json.Marshal(passwd_data)
		// if err != nil{
		// 	ctx.JSON(200, gin.H{"msg": "Null"})
		// 	return
		// }
		var result []map[string]interface{}
		for _, item := range important_data {
			tempMap := map[string]interface{}{
				"key_word":     item.Keyword,
				"connect_user": item.Username,
			}
			result = append(result, tempMap)
		}
		ctx.JSON(200, result)
	})
	r.POST("/MessageManage/Search", func(ctx *gin.Context) {
		//通过关联用户以及应用名字，查询密码
		var requestMap struct {
			Key_word     string `json:"key_word"`
			Connect_user string `json:"connect_user"`
		}
		err := ctx.ShouldBind(&requestMap)
		if err != nil {
			fmt.Println("解析请求体失败:", err)
			ctx.String(http.StatusNotFound, "绑定form失败")
		}
		if requestMap.Key_word != "" && requestMap.Connect_user != "" {
			//两者都有，因此只展示一个
			_, flag := middlebox.Get_keyword_name2key_important(requestMap.Key_word, requestMap.Connect_user)
			if flag != 1 {
				ctx.String(http.StatusBadRequest, "该关联用户未存储该应用密码")
				return
			}
		} else if requestMap.Key_word == "" && requestMap.Connect_user == "" {
			ctx.String(http.StatusBadRequest, "必须输入应用名或关联用户名")
			return
		} else if requestMap.Key_word != "" {
			//var users []string
			users, flag := middlebox.Get_keyword2user_important(requestMap.Key_word)
			if flag != 1 || len(users) == 0 {
				ctx.String(http.StatusBadRequest, "无法查到有关联用户存储着该关键信息")
				return
			}
			var result []map[string]interface{}
			for _, item := range users {
				tempMap := map[string]interface{}{
					"key_word":     requestMap.Key_word,
					"connect_user": item,
				}
				result = append(result, tempMap)
			}
			ctx.JSON(200, result)
		} else {
			keywords, flag := middlebox.Get_user2keyword_important(requestMap.Connect_user)
			if flag != 1 || len(keywords) == 0 {
				ctx.String(http.StatusBadRequest, "无法查到该关联用户存储着任何关键信息")
				return
			}
			var result []map[string]interface{}
			for _, item := range keywords {
				tempMap := map[string]interface{}{
					"key_word":     item,
					"connect_user": requestMap.Connect_user,
				}
				result = append(result, tempMap)
			}
			ctx.JSON(200, result)
		}
	})
	r.POST("/MessageManage/Delete", func(ctx *gin.Context) {
		//通过关联用户以及应用名字，删除该条密码
		var requestMap struct {
			Key_word     string `json:"key_word"`
			Connect_user string `json:"connect_user"`
		}
		err := ctx.ShouldBind(&requestMap)
		if err != nil {
			fmt.Println("解析请求体失败:", err)
			ctx.String(http.StatusNotFound, "绑定form失败")
		}
		if requestMap.Key_word == "" || requestMap.Connect_user == "" {
			ctx.String(http.StatusBadRequest, "信息不全")
			return
		}
		flag := middlebox.Delete_single_important(requestMap.Key_word, requestMap.Connect_user)
		if flag != 1 {
			ctx.String(http.StatusBadRequest, "未删除")
			return
		}
		ctx.JSON(200, gin.H{"msg": "successfully added."})
	})
	// User
	r.POST("/User/ShowPublicKey", func(ctx *gin.Context) {
		//查询自己的公钥按钮
		result, flag := middlebox.Show_myself_public_key()
		if flag != 1 {
			ctx.String(http.StatusBadRequest, "公钥不存在")
			return
		}
		ctx.JSON(200, result)
	})
	r.POST("/User/ShowOtherPublicKey", func(ctx *gin.Context) {
		//查询自己拥有关联用户的公钥
		var requestMap struct {
			Mode           int      `json:"mode"`
			Selecter_users []string `json:"selecter_users"`
		}
		err := ctx.ShouldBind(&requestMap)
		if err != nil {
			fmt.Println("解析请求体失败:", err)
			ctx.String(http.StatusNotFound, "绑定form失败")
		}
		var result []map[string]interface{}
		if requestMap.Mode == 0 {
			//表示正常的选择过程，此时前端传回的用户有几个，就差几个
			for _, item := range requestMap.Selecter_users {
				temp, flag := middlebox.Get_single_public_key(item)
				if flag == 1 {
					tempMap := map[string]interface{}{
						"public_key":   temp,
						"connect_user": item,
					}
					result = append(result, tempMap)
				}
			}
		} else {
			//表示查询全部的用户公钥
			namelist, keylist, flag := middlebox.Get_all_public_key()
			if flag != 1 {
				ctx.String(http.StatusBadRequest, "查询出错")
				return
			}
			for index, _ := range namelist {
				tempMap := map[string]interface{}{
					"public_key":   keylist[index],
					"connect_user": namelist[index],
				}
				result = append(result, tempMap)
			}
		}
		ctx.JSON(200, result)
	})

	// r.Run()
	r.Run("127.0.0.1:8090")
}
