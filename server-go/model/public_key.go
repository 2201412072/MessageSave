package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Public_keys struct {
	Username   string `gorm:"primaryKey;size:100;references:userpassword:user"`
	Public_key []byte `gorm:"type:longblob"`
}

func AddPublicKey(user string, public_key []byte) int {
	temp := Public_keys{Username: user, Public_key: public_key}
	result := database.Create(&temp)
	if result.Error != nil {
		return 0
	}
	return 1
}

func GetPublicKeyByUser(user string) ([]byte, int) {
	var temp Public_keys
	err := database.Where("Username=? ", user).First(&temp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return make([]byte, 0), 0
		} else {
			fmt.Println("Error:", err)
			return make([]byte, 0), 0
		}
	} else {
		return temp.Public_key, 1
	}
}

func GetPublicKey() ([]Public_keys, int) {
	var temps []Public_keys
	database.Find(&temps)
	return temps, 1
}

func DeletePublicKey(user string) int {
	result := database.Where("Username=?", user).Delete(&Public_keys{})
	if result.Error != nil {
		return 0
	}
	return 1
}

// func UpdatePublicKey(user string, public_key []byte)
