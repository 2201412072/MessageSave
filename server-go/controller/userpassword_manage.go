package controller

import (
	"server-go/model"
	"server-go/model/modelview"

	"github.com/gin-gonic/gin"
)

func GetPassword(ctx *gin.Context) {
	// 解析表单输入
	var requestMap struct {
		User string `json:"connect_user"`
	}
	ctx.ShouldBind(&requestMap)
	var ans []model.UserPassword
	if requestMap.User == "" { // 获取所有用户的用户密码
		ans, _ = model.GetAllUserPassword()
	} else {
		temp, _ := model.GetUserPassword(requestMap.User)
		ans = append(ans, temp)
	}
	result := make([]modelview.Password, len(ans))
	for i, v := range ans {
		result[i].User = v.User
		result[i].Password = v.Password
	}
	// 回复前端
	ctx.JSON(200, result)
}

func DeletePassword(ctx *gin.Context) {
	// 解析表单输入
	var requestMap struct {
		User string `json:"connect_user"`
	}
	ctx.ShouldBind(&requestMap)
	// 删除公钥
	model.DeleteUserPassword(requestMap.User)
	// 回复前端
	ctx.JSON(200, gin.H{"msg": "Delete over."})
}
