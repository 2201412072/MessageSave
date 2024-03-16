package controller

import (
	"client-go/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册
func Register(ctx *gin.Context) {
	//向服务器发送用户名和密码
	var requestMap struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := ctx.ShouldBind(&requestMap)
	if err != nil {
		fmt.Println("解析请求体失败:", err)
		ctx.String(http.StatusNotFound, "绑定form失败")
	}
	var sendMessage model.Message
	sendMessage.SrcUser = requestMap.Username
	sendMessage.DstUser = "server"
	sendMessage.KeyWord = requestMap.Password
	sendMessage.Operate = "Register2Server"
	PostMessage(sendMessage)
	response := <-ResponseChan
	switch response {
	case "RegisterError1":
		ctx.JSON(401, gin.H{"msg": "用户名已存在", "result": 0})
	case "RegisterPass":
		ctx.JSON(200, gin.H{"msg": "注册成功", "result": 1})
	default:
		ctx.JSON(401, gin.H{"msg": "发生错误", "result": 0})
	}
}

// 登录
func Login(ctx *gin.Context) {
	//向服务器发送用户名和密码
	var requestMap struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := ctx.ShouldBind(&requestMap)
	if err != nil {
		fmt.Println("解析请求体失败:", err)
		ctx.String(http.StatusNotFound, "绑定form失败")
	}
	var sendMessage model.Message
	sendMessage.SrcUser = requestMap.Username
	sendMessage.DstUser = "server"
	sendMessage.KeyWord = requestMap.Password
	sendMessage.Operate = "Login2Server"
	// login 使用socket通信，限制了一台机器只能有一个用户
	PostMessage(sendMessage)
	response := <-ResponseChan
	fmt.Println("response:", response)
	switch response {
	case "LoginError1":
		ctx.JSON(401, gin.H{"msg": "用户名或密码错误", "result": 0})
	case "LoginError2":
		ctx.JSON(401, gin.H{"msg": "用户名或密码错误", "result": 0})
	case "LoginPass":
		ctx.JSON(200, gin.H{"msg": "登录成功", "result": 1, "whatpage": "client"})
	default:
		ctx.JSON(401, gin.H{"msg": "发生错误", "result": 0})
	}
}
