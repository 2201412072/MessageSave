package router

import (
	"server-go/controller"
	"server-go/util"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(util.CORSHandler()) // 设置全局跨域访问

	r.POST("/ServerHomePage/MessageManage/Message-Search", controller.GetMessage)
	r.POST("/ServerHomePage/MessageManage/Message-Delete", controller.DeleteMessage)
	r.POST("/ServerHomePage/MessageManage/Message-SearchAll", controller.GetMessage)

	r.POST("/ServerHomePage/PublicKeyManage/PublicKey-Search", controller.GetPublicKey)
	r.POST("/ServerHomePage/PublicKeyManage/PublicKey-Delete", controller.DeletePublicKey)
	r.POST("/ClientHomePage/ PublicKeyManage/PublicKey-SearchAll", controller.GetPublicKey)

	r.POST("/ServerHomePage/PasswordManage/Password-Search", controller.GetPassword)
	r.POST("/ServerHomePage/PasswordManage/Password-Delete", controller.DeletePassword)
	r.POST("/ServerHomePage/PasswordManage/Password-SearchAll", controller.GetPassword)
	return r
}
