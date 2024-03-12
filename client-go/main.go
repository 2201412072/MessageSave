package main

import (
	"client-go/controller"
	"client-go/router"
)

func main() {
	//controller.ResponseChan = make(chan string, 1) //用于子协程传递回应消息
	// 导入配置
	// 导入数据库
	controller.Init()
	// 设置路由
	go controller.InitProxy()
	r := router.SetupRouter()
	r.Run("127.0.0.1:8090")

}
