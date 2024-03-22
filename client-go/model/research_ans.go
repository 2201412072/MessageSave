package model

type ResearchAns struct {
	SrcUser     string `gorm:"primaryKey;column:user_src;foreignKey:username;references:public_keys:username"`
	DstUser     string `gorm:"primaryKey;column:user_dst;foreignKey:username;references:public_keys:username"`
	Application string `gorm:"primaryKey;column:application;foreignKey:application;references:password:application"`
	Password    string
	Stage       string
}

// 根据对应值检索结果，输入参数可能存在空值
// 但是我没考虑空值，因为其他我觉得也要考虑，目前都没有考虑，需要后续修改
func GetResearchAns(src_user string, dst_user string, application string, stage string) ([]ResearchAns, int) {
	var temps []ResearchAns
	result := database.Where("user_src=? AND user_dst=? AND application=? AND Stage=?", src_user, dst_user, application, stage).Find(&temps)
	if result.Error != nil {
		return make([]ResearchAns, 0), 0
	} else {
		return temps, 1
	}
}

func GetAllResearchAns() ([]ResearchAns, int) {
	var temps []ResearchAns
	result := database.Find(&temps)
	if result.Error != nil {
		return make([]ResearchAns, 0), 0
	} else {
		return temps, 1
	}
}

func GetResearchAnsByDst(dst_user string) ([]ResearchAns, int) {
	var temps []ResearchAns
	result := database.Where("user_dst=?", dst_user).Find(&temps)
	if result.Error != nil {
		return make([]ResearchAns, 0), 0
	} else {
		return temps, 1
	}
}

func GetResearchAnsByApp(application string) ([]ResearchAns, int) {
	var temps []ResearchAns
	result := database.Where("application=?", application).Find(&temps)
	if result.Error != nil {
		return make([]ResearchAns, 0), 0
	} else {
		return temps, 1
	}
}

func AddResearchAns(src_user string, dst_user string, application string, stage string, password string) int {
	temp := ResearchAns{SrcUser: src_user, DstUser: dst_user, Application: application, Stage: stage, Password: password}
	result := database.Create(&temp)
	if result.Error != nil {
		return 0
	}
	return 1
}

func DeleteResearchAns(src_user string, dst_user string, application string) int {
	result := database.Where("user_src=? AND user_dst=? AND application=?", src_user, dst_user, application).Delete(&ResearchAns{})
	if result.Error != nil {
		return 0
	}
	return 1
}

func DeleteResearchAnsByApp(application string) int {
	result := database.Where("application=?", application).Delete(&ResearchAns{})
	if result.Error != nil {
		return 0
	}
	return 1
}

func ChangeResearchAns(src_user string, dst_user string, application string, stage string, password string) int {
	//该函数用于修改状态，它表示一个src_user,dst_user_application的键值的状态被修改为stage，密码被修改为password
	result := database.Model(&ResearchAns{}).Where("user_src = ? AND user_dst=? AND application=?", src_user, dst_user, application).Updates(map[string]interface{}{"Stage": stage, "Password": password})
	if result.Error != nil {
		return 0
	}
	return 1
}
