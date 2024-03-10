package model

type Message struct {
	SrcUser string `gorm:"primaryKey;column:user_src;references:userpassword:user"`
	DstUser string `gorm:"primaryKey;column:user_dst;references:userpassword:user"`
	KeyWord string `gorm:"primaryKey;column:keyword;references:password:application"`
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

func GetMessageBySrcUser(src_user string) ([]Message, int) {
	var temp []Message
	database.Where("user_src=? ", src_user).Find(&temp)
	return temp, 1
}

func GetMessageByDstUser(dst_user string) ([]Message, int) {
	var temp []Message
	database.Where("user_dst=? ", dst_user).Find(&temp)
	return temp, 1
}

func GetMessageByApplication(application string) ([]Message, int) {
	var temp []Message
	database.Where("keyword=? ", application).Find(&temp)
	return temp, 1
}

func GetMessageByOperate(operate string) ([]Message, int) {
	var temp []Message
	database.Where("Operate=? ", operate).Find(&temp)
	return temp, 1
}

func AddMessage(msg Message) int {
	result := database.Create(&msg)
	if result.Error != nil {
		return 0
	}
	return 1
}

func DeleteMessage(src_user string, dst_user string, application string) int {
	result := database.Where("user_src=? AND user_dst=? AND keyword=?", src_user, dst_user, application).Delete(&Message{})
	if result.Error != nil {
		return 0
	}
	return 1
}
