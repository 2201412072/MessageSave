package controller

import (
	"server-go/model"

	"github.com/gin-gonic/gin"
)

func GetPassword(ctx *gin.Context) {
	// 解析表单输入
	var requestMap string
	ctx.ShouldBind(&requestMap)
	var ans []model.UserPassword
	if requestMap == "" { // 获取所有用户的用户密码
		ans, _ = model.GetAllUserPassword()
	} else {
		temp, _ := model.GetUserPassword(requestMap)
		ans = append(ans, temp)
	}
	// 回复前端
	ctx.JSON(200, ans)
}

func DeletePassword(ctx *gin.Context) {
	// 解析表单输入
	var requestMap string
	ctx.ShouldBind(&requestMap)
	// 删除公钥
	model.DeleteUserPassword(requestMap)
	// 回复前端
	ctx.JSON(200, gin.H{"msg": "Delete over."})
}
