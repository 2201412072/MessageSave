package controller

import (
	"server-go/model"
)

var server_addr string = "127.0.0.1"
var port string = "8091"

// // 修改数据库存储路径
// func Change_database_path(db_path string) int {
// 	database_path = db_path
// 	return 1
// }

func Init() {
	// 加载配置文件
	// 加载数据库
	model.SetupDB()
}
