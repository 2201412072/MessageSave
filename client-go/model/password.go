package model

import (
	"client-go/util"
)

type Password struct {
	Application string `json:"app";gorm:"primaryKey;column:application"`
	Username    string `json:"connect_user";gorm:"primaryKey;column:username;foreignKey:username;references:public_keys:username"`
	Saved_key   []byte `json:"password";gorm:"type:longblob"`
	Single_key  string `gorm:"size:10"`
}

func AddPassword(application string, username string, saved_key []byte, single_key string) int {
	temp := Password{Username: username, Application: application, Saved_key: saved_key, Single_key: single_key}
	result := database.Create(&temp)
	if result.Error != nil {
		return 0
	}
	return 1
}

func DeletePassword(application string, username string) int {
	result := database.Where("username=? AND application=?", username, application).Delete(&Password{})
	if result.Error != nil {
		return 0
	}
	return 1
}

func UpdatePassword(application string, username string, saved_key []byte, single_key string) int {
	result := database.Model(&Password{}).Where("username=? AND application=?", username, application).Update("Saved_key", saved_key)
	if result.Error != nil {
		return 0
	}
	return 1
}

func GetPasswordString(username string, application string) (string, int) {
	//该函数用于获取数据库中的Save-key，但是它需要转成string后返回
	var temp Password
	database.Where("username=? AND application=?", username, application).First(&temp)
	str, flag := util.Bytes2base_string(temp.Saved_key)
	if flag != 1 {
		return "", 0
	}
	return str, 1
}

func GetPassword(username string, application string) (Password, int) {
	var temps []Password
	result := database.Where("application=? AND username=?", application, username).Find(&temps)
	if result.Error != nil {
		return Password{}, 0
	} else {
		if len(temps) != 1 {
			return Password{}, 0
		} else {
			return temps[0], 1
		}
	}
}

func GetPasswordByApp(application string) ([]Password, int) {
	var temps []Password
	result := database.Where("application=?", application).Find(&temps)
	if result.Error != nil {
		return temps, 0
	} else {
		return temps, 1
	}
}

func GetPasswordByUser(username string) ([]Password, int) {
	var temps []Password
	result := database.Where("username=?", username).Find(&temps)
	if result.Error != nil {
		return temps, 0
	} else {
		return temps, 1
	}
}

func GetALLPassword() ([]Password, int) {
	var temps []Password
	result := database.Find(&temps)
	if result.Error != nil {
		return temps, 0
	} else {
		return temps, 1
	}
}

//我不知道下面两个函数是在干嘛

func GetUserByPasswordInPassword(username string) int {
	return 1
}

func GetUserByKeyWordInPassword(key_word string) int {
	return 1
}
