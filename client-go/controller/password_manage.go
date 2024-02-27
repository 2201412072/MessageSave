package controller

import (
	"client-go/model"
	"client-go/model/modelview"

	"github.com/gin-gonic/gin"
)

func SearchPasssword(ctx *gin.Context) {
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
	var result []map[string]interface{}
	for _, item := range passwd2_data {
		tempMap := map[string]interface{}{
			"key_word":     item.Application,
			"connect_user": item.Username,
			// "password2":    item.Saved_key,
		}
		result = append(result, tempMap)
	}
	ctx.JSON(200, result)
}

// 删除指定密码
func DeletePasssword(ctx *gin.Context) {
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
