package model

type ResearchAns struct {
	SrcUser     string `gorm:"primaryKey;column:user_src;foreignKey:username;references:public_keys:username"`
	DstUser     string `gorm:"primaryKey;column:user_dst;foreignKey:username;references:public_keys:username"`
	Application string `gorm:"primaryKey;column:application;foreignKey:application;references:password:application"`
	Stage       string
}

// 根据对应值检索结果，输入参数可能存在空值
func GetResearchAns(src_user string, dst_user string, application string, stage string) ([]ResearchAns, int) {
	var temp []ResearchAns
	return temp, 1
}

func AddResearchAns(src_user string, dst_user string, application string, stage string) int {
	return 1
}

func DeleteResearchAns(src_user string, dst_user string, application string) int {
	return 1
}

func ChangeResearchAns(src_user string, dst_user string, application string, stage string, password string) int {
	//该函数用于修改状态，它表示一个src_user,dst_user_application的键值的状态被修改为stage，密码被修改为password
	return 1
}
