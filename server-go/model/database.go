package model

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB
var err error

// 启动数据库
func SetupDB() {
	user := "message_save"
	password := "message_save"
	host := "localhost"
	port := "3306"
	dbname := "MessageSaveDB"
	charset := "utf8"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", user, password, host, port, dbname, charset)

	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("myDB open failed!", err)
	}

	// 数据库初始化表（默认数据库已建立）
	DB_init_table()
}

// 初始化数据库各张表
func DB_init_table() {
	database.AutoMigrate(&Message{})
	database.AutoMigrate(&UserPassword{})
	database.AutoMigrate(&Public_keys{})
	database.AutoMigrate(&ResearchAns{})
}
