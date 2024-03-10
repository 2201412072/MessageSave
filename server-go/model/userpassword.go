package model

import (
	"fmt"

	"gorm.io/gorm"
)

//该表用来记录登录的用户名和密码

type UserPassword struct {
	User     string `gorm:"primaryKey;column:user"`
	Password string
}

func GetAllUserPassword() ([]UserPassword, int) {
	var temps []UserPassword
	database.Find(&temps)
	return temps, 1
}

func GetUserPassword(user string) (UserPassword, int) {
	var temp UserPassword
	err := database.Where("User=? ", user).First(&temp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return temp, 0
		} else {
			fmt.Println("Error:", err)
			return temp, 0
		}
	} else {
		return temp, 1
	}
}

func AddUserPassword(user string, password string) int {
	temp := UserPassword{User: user, Password: password}
	result := database.Create(&temp)
	if result.Error != nil {
		return 0
	}
	return 1
}

func DeleteUserPassword(user string) int {
	result := database.Where("User=?", user).Delete(&UserPassword{})
	if result.Error != nil {
		return 0
	}
	return 1
}
