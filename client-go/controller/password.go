package controller

import (
	"client-go/model"
	"client-go/model/modelview"
	"client-go/util"

	"github.com/gin-gonic/gin"
)

// 存储密码
func SavePassword(ctx *gin.Context) {
	// 解析表单输入
	var requestMap modelview.Password
	ctx.ShouldBind(&requestMap)
	user := requestMap.User
	passwd := requestMap.Password
	application := requestMap.Application
	// 对密码进行本地加密
	other_public_key, _ := GetPublicKeyByUser(user) // proxy.GetPublicKey(user)
	passwd1, _ := util.EncryptUTFString(passwd, public_key)
	passwd2, _ = util.Block_encrypt(passwd1, other_public_key)

	single_key := ""
	// 存储至数据库
	model.AddPassword(application, user, passwd2, single_key)
	// 回复前端
	ctx.JSON(200, gin.H{"msg": "password saved."})
}

// 使用密码
func UsePassword(ctx *gin.Context) {
	// 解析表单数据
	var requestMap modelview.Password
	ctx.ShouldBind(&requestMap)
	user := requestMap.User
	application := requestMap.Application
	// 获取加密密码
	passwd2, _ := model.GetPasswordString(user, application)
	// 消息代理向服务器发送关联用户解密请求
	RequireDecrypt(user, application, passwd2)
	// 回复前端
	ctx.JSON(200, gin.H{"msg": "已向对方发送解密请求。"})
}

// 解密一次加密的密文，返回真实密码
func DecryptPassword1(message string) string {
	return ""
}
