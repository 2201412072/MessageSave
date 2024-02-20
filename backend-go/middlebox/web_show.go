package middlebox

import (
	"mydatabase"

	"gorm.io/gorm"
)

// var MyDB *gorm.DB
// var database_path string = "database/mysql"

type passwd_view struct {
	key_word         string `gorm:"primaryKey;column:key_word"`
	connect_user     string `column:connect_user;foreignKey:username;references:public_keys:username"`
	encrypted_passwd string `gorm:"type:longblob"`
	decrypted_passwd string `gorm:"size:100"`
}

func Init_web_show_table() int {
	MyDB, flag := mydatabase.Db_load()
	if flag != 1 {
		return 0
	}

	// 配置前端密码视图
	if !MyDB.Migrator().HasTable(&passwd_view{}) {
		/*
			// CREATE VIEW `user_view` AS SELECT * FROM `users` WHERE age > 20
			q := DB.Model(&User{}).Where("age > ?", 20)
			DB.Debug().Migrator().CreateView("user_view", gorm.ViewOption{Query: q})
		*/
		// create view `web_passwd` as select application,username,saved_key,'' from password
		passwd_view_q := MyDB.Model(&passwd_view{}) // ?
		err := MyDB.Migrator().CreateView("passwd_view", gorm.ViewOption{Query: passwd_view_q})
		if err != nil {
			return 0
		}
	}

	// 配置前端重要信息视图
	return 1
}
