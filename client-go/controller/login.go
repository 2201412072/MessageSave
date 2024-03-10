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
		username string `json:"username"`
		password string `json:"password"`
	}
	err := ctx.ShouldBind(&requestMap)
	if err != nil {
		fmt.Println("解析请求体失败:", err)
		ctx.String(http.StatusNotFound, "绑定form失败")
	}
	var sendMessage model.Message
	sendMessage.SrcUser = requestMap.username
	sendMessage.DstUser = "server"
	sendMessage.KeyWord = requestMap.password
	sendMessage.Operate = "Register2Server"
	PostMessage(sendMessage)
	ctx.JSON(200, gin.H{"msg": "message sent successfully"})
}

// 登录
func Login(ctx *gin.Context) {
	//向服务器发送用户名和密码
	var requestMap struct {
		username string `json:"username"`
		password string `json:"password"`
	}
	err := ctx.ShouldBind(&requestMap)
	if err != nil {
		fmt.Println("解析请求体失败:", err)
		ctx.String(http.StatusNotFound, "绑定form失败")
	}
	var sendMessage model.Message
	sendMessage.SrcUser = requestMap.username
	sendMessage.DstUser = "server"
	sendMessage.KeyWord = requestMap.password
	sendMessage.Operate = "Login2Server"
	PostMessage(sendMessage)
	ctx.JSON(200, gin.H{"msg": "message sent successfully"})
}
