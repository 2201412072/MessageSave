package router

import (
	"client-go/configs"
	"client-go/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 注册
	r.POST("/Register", controller.Register)
	// 登录
	r.POST("/Login", controller.Login)
	// 设置
	r.GET("/Setting", configs.Setting)
	// 修改设置
	// 查询公钥
	r.POST("/ClientHomePag/Config/PublicKey-Search", controller.GetPublicKey)
	// 删除公钥
	r.POST("/ClientHomePag/Config/PublicKey-Delete", controller.DeletePublicKey)
	// 密码存储
	r.POST("/ClientHomePage/AddPassword", controller.SavePassword)
	// 密码使用
	r.POST("/ClientHomePage/SearchPassword", controller.UsePassword)
	// // 密码管理（跟密码管理查询放一块吧）
	// r.POST("/PasswordManage",)
	// 密码管理查询
	r.POST("/ClientHomePage/PasswordManage/SearchPassword", controller.SearchPasssword)
	// 密码管理添加密码
	r.POST("/ClientHomePage/PasswordManage/AddPassword", controller.SavePassword)
	// 密码管理查询密码
	r.POST("/ClientHomePage/PasswordManage/Password-Search", controller.SearchPasssword)
	// 密码管理删除
	r.POST("/ClientHomePage/PasswordManage/Password-Delete", controller.DeletePasssword)
	// // 密码管理结果查询
	// r.POST("/ClientHomePage/PasswordManage/Result-Search", controller.GetMessage)
	// // 密码管理结果删除
	// r.POST("/ClientHomePag/PasswordManage/Result-Delete", controller.DeleteMessage)
	// 消息管理
	// 消息查询
	r.POST("/ClientHomePage/Announcement/Request-Search", controller.GetMessage)
	// 消息删除
	r.POST("/ClientHomePage/Announcement/Request-Delete", controller.DeleteMessage)
	// 解密请求同意
	r.POST("/ClientHomePage/Announcement/Request-Agree", controller.AgreeMessage)
	// 解密请求拒绝
	r.POST("/ClientHomePage/Announcement/Request-Disagree", controller.DisagreeMessage)

	return r
}
