package model

type Message struct {
	SrcUser string `gorm:"primaryKey;column:user_src;foreignKey:username;references:public_keys:username"`
	DstUser string `gorm:"primaryKey;column:user_dst;foreignKey:username;references:public_keys:username"`
	KeyWord string `gorm:"primaryKey;column:keyword;foreignKey:keyword;references:password:application"`
	// Time         time.Time `gorm:"column:mytime"` //primaryKey;
	Operate string
	Params  string
}

func GetMessages() ([]Message, int) {
	var temp []Message
	database.Find(&temp)
	return temp, 1
}

func GetMessage(src_user string, dst_user string, key_word string) (Message, int) {
	var temp Message
	database.Where("SrcUser=? AND DstUser=? AND KeyWord=?", src_user, dst_user, key_word).First(&temp)
	return temp, 1
}

func GetMessageBySrcUser(src_user string) int {
	return 1
}

func GetMessageByDstUser(dst_user string) int {
	return 1
}

func GetMessafeByApplication(application string) int {
	return 1
}

func GetMessageByOperate(operate string) int {
	return 1
}

func AddMessage(msg Message) int {
	return 1
}

func DeleteMessage(src_user string, dst_user string, application string) int {
	return 1
}