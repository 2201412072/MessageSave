package main

import (
	"server-go/controller"
	"server-go/router"
)

func main() {
	// 导入配置
	// 导入数据库
	controller.Init()
	// 设置路由
	r := router.SetupRouter()
	r.Run("127.0.0.1:8090")
	go controller.InitProxy()
}
