package model

type Public_keys struct {
	Username   string `gorm:"primaryKey;size:100"`
	Public_key []byte `gorm:"type:longblob"`
}

func AddPublicKey(user string, public_key []byte) int {
	return 1
}

func GetPublicKeyByUser(user string) ([]byte, int) {
	return make([]byte, 0), 1
}

func GetPublicKey() ([][]byte, int) {
	rows := 1
	temp := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		temp[i] = make([]byte, 0)
	}
	return temp, 1
}

func DeletePublicKey(user string) int {
	return 1
}

// func UpdatePublicKey(user string, public_key []byte)
