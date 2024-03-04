package controller

import (
	"client-go/model"

	"github.com/gin-gonic/gin"
)

// 获取查询结果
func GetResearchAns(ctx *gin.Context) {
	// 解析表单输入
	var requestMap model.ResearchAns
	ctx.ShouldBind(&requestMap)
	dst_user := requestMap.DstUser
	application := requestMap.Application
	stage := requestMap.Stage
	// 查询结果
	ans, _ := model.GetResearchAns(username, dst_user, application, stage)
	// 回复前端
	ctx.JSON(200, ans)
}

// 删除查询结果
func DeleteResearchAns(ctx *gin.Context) {
	// 解析表单输入
	var requestMap model.ResearchAns
	ctx.ShouldBind(&requestMap)
	dst_user := requestMap.DstUser
	application := requestMap.Application
	// 删除查询结果
	model.DeleteResearchAns(username, dst_user, application)
	// 回复前端
	ctx.JSON(200, gin.H{"msg": "delete over."})
}
