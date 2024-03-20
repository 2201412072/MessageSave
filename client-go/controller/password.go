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
	// user := requestMap.Connect_user
	// passwd := requestMap.Password
	// application := requestMap.App
	// 对密码进行本地加密
	other_public_key, flag := GetPublicKeyByUser(user) // proxy.GetPublicKey(user)
	//flag=0出bug了，=1为正常，=2表示服务器没找到该用户
	switch flag {
	case 0:
		ctx.JSON(401, gin.H{"msg": "bug!!!"})
	case 1:
		{
			passwd1, _ := util.EncryptUTFString(passwd, Public_key)
			passwd2, _ := util.Block_encrypt(passwd1, other_public_key)

			single_key := ""
			// 存储至数据库
			flag := model.AddPassword(application, user, passwd2, single_key)
			if flag == 0 {
				//添加密码出现了故障，应该是已经有该条数据了，所以需要先删除
				ctx.JSON(401, gin.H{"msg": "已经通过" + user + "保存了" + application + "的密码，请先删除"})
				return
			}
			//将该加密信息发送给服务器，然后服务器代为转发给对应的客户端
			var message model.Message
			message.DstUser = user
			message.Operate = "EncryptAnnocement2Server"
			message.KeyWord = application
			PostMessage(message)
			//接收服务器返回的信息，但是我们不关心服务器添加没添加这个消息
			_ = <-ResponseChan
			// 回复前端
			ctx.JSON(200, gin.H{"msg": "password saved."})
		}

	case 2:
		ctx.JSON(401, gin.H{"msg": "服务器没找到该用户"})
	}
}

// 使用密码
func UsePassword(ctx *gin.Context) {
	// 解析表单数据
	var requestMap modelview.Password
	ctx.ShouldBind(&requestMap)
	user := requestMap.User
	application := requestMap.Application
	// user := requestMap.Connect_user
	// application := requestMap.App
	// 获取加密密码
	passwd2, _ := model.GetPasswordString(user, application)
	// 消息代理向服务器发送关联用户解密请求
	RequireDecrypt(user, application, passwd2)
	//接收服务器返回的信息
	response := <-ResponseChan
	if response == "AddMessageError1" {
		ctx.JSON(401, gin.H{"msg": "服务器未转发，请重新发送"})
		return
	} else if response == "AddMessagePass" {
		//将该消息存入research_ans表中，方便查找
		model.AddResearchAns(username, user, application, "hasn't completed", "")

		// 回复前端
		ctx.JSON(200, gin.H{"msg": "已向对方发送解密请求。"})
	}

}

// 解密一次加密的密文，返回真实密码
func DecryptPassword1(message string) string {
	//这里假定服务器向客户端传送的依旧像第一版那样，是base64字符串。（但事实上不用，直接传二进制值就行了）
	temp_bytes_data, flag := util.Base_string2bytes(message)
	if flag != 1 {
		return ""
	}
	bytes_data, _ := util.Block_decrypt(temp_bytes_data, Private_key)
	string_data, flag := util.Base_bytes2utf_string(bytes_data)
	if flag != 1 {
		return ""
	}
	return string_data

}
