package controller

import (
	"server-go/model"
	"server-go/model/modelview"
	"server-go/util"

	"github.com/gin-gonic/gin"
)

// 获得储存的用户公钥
func GetPublicKey(ctx *gin.Context) {
	// 解析表单输入
	var requestMap struct {
		User string `json:"connect_user"`
	}
	ctx.ShouldBind(&requestMap)
	// 获取公钥们并转换为字符串
	var ans []modelview.PublicKeyResponse
	if requestMap.User == "" { // 获取所有用户的公钥
		public_keys, _ := model.GetPublicKey()
		ans = make([]modelview.PublicKeyResponse, len(public_keys))
		for i := 0; i < len(public_keys); i++ {
			// ans[i],_ = util.PublicKey_from_bytes(public_keys[i])
			ans[i].PublicKey, _ = util.Bytes2base_string(public_keys[i].Public_key)
			ans[i].User = public_keys[i].Username
		}
	} else {
		var public_key []byte
		public_key, _ = model.GetPublicKeyByUser(requestMap.User)
		ans = make([]modelview.PublicKeyResponse, 1)
		ans[0].PublicKey, _ = util.Bytes2base_string(public_key)
		ans[0].User = requestMap.User
	}
	// 回复前端
	ctx.JSON(200, ans)
}

// 删除指定的用户公钥
func DeletePublicKey(ctx *gin.Context) {
	// 解析表单输入
	var requestMap struct {
		User string `json:"connect_user"`
	}
	ctx.ShouldBind(&requestMap)
	// 删除公钥
	model.DeletePublicKey(requestMap.User)
	// 回复前端
	ctx.JSON(200, gin.H{"msg": "Delete over."})
}
