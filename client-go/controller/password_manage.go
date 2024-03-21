package controller

import (
	"client-go/model"
	"client-go/model/modelview"

	"github.com/gin-gonic/gin"
)

func AggUsers(data []model.Password) map[string][]string {
	result := make(map[string][]string)
	for _, item := range data {
		result[item.Application] = append(result[item.Application], item.Username)
	}
	return result
}

func SearchPassword(ctx *gin.Context) {
	// 搜索管理的密码
	// 解析表单数据
	var requestMap model.Password
	ctx.ShouldBind(&requestMap)
	user := requestMap.Username
	application := requestMap.Application
	// 获取数据库对应密码
	var passwd2_data []model.Password
	if user == "" && application == "" {
		passwd2_data, _ = model.GetALLPassword()
	} else if user != "" && application == "" {
		passwd2_data, _ = model.GetPasswordByUser(user)
	} else if user == "" && application != "" {
		passwd2_data, _ = model.GetPasswordByApp(application)
	} else {
		passwd2_data := make([]model.Password, 1)
		passwd2_data[0], _ = model.GetPassword(user, application)
	}
	// 回复前端
	// var result []map[string]interface{}
	// for _, item := range passwd2_data {
	// 	tempMap := map[string]interface{}{
	// 		"app":          item.Application,
	// 		"connect_user": item.Username,
	// 		// "password2":    item.Saved_key,
	// 	}
	// 	result = append(result, tempMap)
	// }
	var result []map[string]interface{}
	temp_data := AggUsers(passwd2_data)
	for key, value := range temp_data {
		tempMap := map[string]interface{}{
			"app":          key,
			"connect_user": value,
		}
		result = append(result, tempMap)
	}
	ctx.JSON(200, result)
}

// 删除指定密码
func DeletePassword(ctx *gin.Context) {
	// 解析表单输入
	var requestMap modelview.Password
	ctx.ShouldBind(&requestMap)
	user := requestMap.User
	application := requestMap.Application
	// 删除数据库中的密码
	model.DeletePassword(application, user)
	// 回复前端
	ctx.JSON(200, gin.H{"msg": "delete over."})
}
