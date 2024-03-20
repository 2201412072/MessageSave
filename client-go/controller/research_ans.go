package controller

import (
	"client-go/model"
	"client-go/model/modelview"

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
	var researchans_data []model.ResearchAns
	if dst_user == "" && application == "" {
		researchans_data, _ = model.GetAllResearchAns()
	} else if dst_user != "" && application == "" {
		researchans_data, _ = model.GetResearchAnsByDst(dst_user)
	} else if dst_user == "" && application != "" {
		researchans_data, _ = model.GetResearchAnsByApp(application)
	} else {
		researchans_data, _ = model.GetResearchAns(username, dst_user, application, stage)
	}
	result := make([]modelview.ResearchAns, len(researchans_data))
	for i, v := range researchans_data {
		result[i].DstUser = v.DstUser
		result[i].Application = v.Application
		result[i].Stage = v.Stage
		result[i].Password, _ = Deal_B2A_message_to_utf(v.Password)
	}
	//ans, _ := model.GetResearchAns(username, dst_user, application, stage)
	// 回复前端
	ctx.JSON(200, result)
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
