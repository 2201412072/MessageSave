package main

import (
	"client-go/router"
)

func main() {
	// 导入配置
	// 导入数据库
	// 设置路由
	r := router.SetupRouter()
	r.Run("127.0.0.1:8090")
}
