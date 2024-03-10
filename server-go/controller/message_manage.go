package controller

import (
	"server-go/model"
	"server-go/model/modelview"

	"github.com/gin-gonic/gin"
)

func GetMessage(ctx *gin.Context) {
	//搜索message，具体而言就是考虑不同参数是否为空的问题
	var requestMap model.Message
	ctx.ShouldBind(&requestMap)
	src_user := requestMap.SrcUser
	dst_user := requestMap.DstUser
	keyword := requestMap.KeyWord
	// 获取数据库对应密码
	var message_data []model.Message
	if src_user == "" && dst_user == "" && keyword == "" {
		message_data, _ = model.GetMessages()
	} else if src_user == "" {
		message_data, _ = model.GetMessageBySrcUser(src_user)
	} else if dst_user == "" {
		message_data, _ = model.GetMessageByDstUser(dst_user)
	} else if keyword == "" {
		message_data, _ = model.GetMessageByApplication(keyword)
	} else {
		message_data := make([]model.Message, 1)
		message_data[0], _ = model.GetMessage(src_user, dst_user, keyword)
	}
	// 回复前端
	var result []map[string]interface{}
	for _, item := range message_data {
		tempMap := map[string]interface{}{
			"src_user":    item.SrcUser,
			"dst_user":    item.DstUser,
			"application": item.KeyWord,
			"operate":     item.Operate,
			"param":       item.Params,
			// "password2":    item.Saved_key,
		}
		result = append(result, tempMap)
	}
	ctx.JSON(200, result)

}

func DeleteMessage(ctx *gin.Context) {
	// 解析表单输入
	var requestMap modelview.Message
	ctx.ShouldBind(&requestMap)
	src_user := requestMap.SrcUser
	dst_user := requestMap.DstUser
	application := requestMap.KeyWord
	// 删除数据库中的密码
	model.DeleteMessage(src_user, dst_user, application)
	// 回复前端
	ctx.JSON(200, gin.H{"msg": "delete over."})
}
