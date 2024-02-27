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
	return temp, 1
}

func GetMessageByUser(user string) int {
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

func DeleteMessage(user string, application string) int {
	return 1
}
