package controller

import (
	"server-go/model"
	"server-go/util"

	"github.com/gin-gonic/gin"
)

// 获得储存的用户公钥
func GetPublicKey(ctx *gin.Context) {
	// 解析表单输入
	var requestMap string
	ctx.ShouldBind(&requestMap)
	// 获取公钥们并转换为字符串
	var ans []string
	if requestMap == "" { // 获取所有用户的公钥
		public_keys, _ := model.GetPublicKey()
		ans = make([]string, len(public_keys))
		for i := 0; i < len(public_keys); i++ {
			// ans[i],_ = util.PublicKey_from_bytes(public_keys[i])
			ans[i], _ = util.Base_bytes2utf_string(public_keys[i])
		}
	} else {
		var public_key []byte
		public_key, _ = model.GetPublicKeyByUser(requestMap)
		ans = make([]string, 1)
		ans[0], _ = util.Base_bytes2utf_string(public_key)
	}
	// 回复前端
	ctx.JSON(200, ans)
}

// 删除指定的用户公钥
func DeletePublicKey(ctx *gin.Context) {
	// 解析表单输入
	var requestMap string
	ctx.ShouldBind(&requestMap)
	// 删除公钥
	model.DeletePublicKey(requestMap)
	// 回复前端
	ctx.JSON(200, gin.H{"msg": "Delete over."})
}
