package model

type Public_keys struct {
	Username   string `gorm:"primaryKey;size:100"`
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
	var temps []Public_keys
	var err int
	var temp []byte
	database.Where("Username = ? ", user).Find(&temps)
	if len(temps) > 1 || len(temps) == 0 {
		err = 0
		temp = make([]byte, 0)
	} else {
		temp = temps[0].Public_key
		err = 1
	}
	return temp, err
}

// func GetPublicKey() ([][]byte, int) {
// 	var temps []Public_keys
// 	database.Find(&temps)
// 	temp := make([][]byte, len(temps))
// 	for i := 0; i < len(temps); i++ {
// 		temp[i] = temps[i].Public_key
// 	}
// 	return temp, 1
// }

func GetPublicKey() ([]Public_keys, int) {
	var temps []Public_keys
	database.Find(&temps)
	return temps, 1
}

func DeletePublicKey(user string) int {
	result := database.Delete(&Public_keys{}, user)
	if result.Error != nil {
		return 0
	}
	return 1
}

// func UpdatePublicKey(user string, public_key []byte)
