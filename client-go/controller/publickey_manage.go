package controller

import (
	"client-go/model"
	"client-go/model/modelview"
	"client-go/util"

	"github.com/gin-gonic/gin"
)

// 获得储存的用户公钥
func GetPublicKey(ctx *gin.Context) {
	// 解析表单输入
	var requestMap modelview.PublicKeyUserRcv
	ctx.ShouldBind(&requestMap)
	// 获取公钥们并转换为字符串
	var ans []modelview.PublicKeyUserSend
	if requestMap.Username == "" { // 获取所有用户的公钥
		public_keys, _ := model.GetPublicKey()
		if len(public_keys) > 0 {
			ans = make([]modelview.PublicKeyUserSend, len(public_keys))
			for i := 0; i < len(public_keys); i++ {
				// ans[i],_ = util.PublicKey_from_bytes(public_keys[i])
				ans[i].Public_key, _ = util.Bytes2base_string(public_keys[i].Public_key)
				ans[i].Connect_user = public_keys[i].Username
			}
		}
	} else {
		var public_key []byte
		public_key, _ = model.GetPublicKeyByUser(requestMap.Username)
		if len(public_key) > 0 {
			ans = make([]modelview.PublicKeyUserSend, 1)
			ans[0].Public_key, _ = util.Bytes2base_string(public_key)
			ans[0].Connect_user = requestMap.Username
		}
	}
	// 回复前端
	ctx.JSON(200, ans)
}

// 删除指定的用户公钥
func DeletePublicKey(ctx *gin.Context) {
	// 解析表单输入
	var requestMap modelview.PublicKeyUserRcv
	ctx.ShouldBind(&requestMap)
	// 删除公钥
	model.DeletePublicKey(requestMap.Username)
	// 回复前端
	ctx.JSON(200, gin.H{"msg": "Delete over."})
}
