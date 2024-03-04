package model

type Password struct {
	Application string `gorm:"primaryKey;column:application"`
	Username    string `gorm:"primaryKey;column:username;foreignKey:username;references:public_keys:username"`
	Saved_key   []byte `gorm:"type:longblob"`
	Single_key  string `gorm:"size:100"`
}

func AddPassword(application string, username string, saved_key string, single_key string) int {
	return 1
}

func DeletePassword(application string, username string) int {
	return 1
}

func UpdatePassword(application string, username string, saved_key string, single_key string) int {
	return 1
}

func GetPasswordString(username string, application string) (string, int) {
	return "", 1
}

func GetPassword(username string, application string) (Password, int) {
	return Password{}, 1
}

func GetPasswordByApp(application string) ([]Password, int) {
	temp := make([]Password, 0)
	return temp, 1
}

func GetPasswordByUser(username string) ([]Password, int) {
	temp := make([]Password, 0)
	return temp, 1
}

func GetALLPassword() ([]Password, int) {
	temp := make([]Password, 0)
	return temp, 1
}

func GetUserByPasswordInPassword(username string) int {
	return 1
}

func GetUserByKeyWordInPassword(key_word string) int {
	return 1
}
