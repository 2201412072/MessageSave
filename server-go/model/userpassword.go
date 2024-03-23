package model

import (
	"fmt"

	"gorm.io/gorm"
)

//该表用来记录登录的用户名和密码,现在还要记录一下该用户是否在线

type UserPassword struct {
	User     string `gorm:"primaryKey;column:user"`
	Password string
	Stage    int
}

//stage表示该用户是否在线，如果在线为1，否则为0

func GetAllUserPassword() ([]UserPassword, int) {
	var temps []UserPassword
	database.Find(&temps)
	return temps, 1
}

func GetUserPassword(user string) (UserPassword, int) {
	var temp UserPassword
	err := database.Where("user=? ", user).First(&temp).Error
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
	temp := UserPassword{User: user, Password: password, Stage: 1}
	result := database.Create(&temp)
	if result.Error != nil {
		return 0
	}
	return 1
}

func DeleteUserPassword(user string) int {
	result := database.Where("user=?", user).Delete(&UserPassword{})
	if result.Error != nil {
		return 0
	}
	return 1
}

func UpdateStageUserPassword(user string, stage int) int {
	//该函数用于调整user的在线状态
	result := database.Model(&UserPassword{}).Where("user = ? ", user).Updates(map[string]interface{}{"Stage": stage})
	if result.Error != nil {
		return 0
	}
	return 1
}
