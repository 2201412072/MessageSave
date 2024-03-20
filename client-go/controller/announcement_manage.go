package controller

import (
	"client-go/model"
	"client-go/model/modelview"

	"github.com/gin-gonic/gin"
)

// 获取未处理的消息，最开始的版本，思路是本地缓存的未处理的消息+服务器消息
// func GetMessage(ctx *gin.Context) {
// 	// 获取本地缓存的未处理的消息
// 	local_messages, _ := model.GetMessages()
// 	// 获取服务器消息并缓存
// 	var net_messages []model.Message
// 	util.ConnSend("GetMessageByUser " + username)
// 	msg_json := util.ConnRecive()
// 	err := json.Unmarshal([]byte(msg_json), &net_messages)
// 	if err != nil {
// 		fmt.Println("Error! ", err)
// 	}
// 	for _, msg := range net_messages { // 缓存
// 		model.AddMessage(msg)
// 	}
// 	// 合并两堆消息
// 	messages := append(local_messages, net_messages...)
// 	// 消息队列
// 	// for _,msg := range messages {
// 	// 	msgQueue.PushBack(msg)
// 	// }
// 	// 回复前端
// 	ctx.JSON(200, messages)
// }

// 获取未处理的消息，最近的版本，思路是：只获取本地缓存消息，至于服务器端的消息，单独设置一个函数、消息格式用来爬取，
// 该函数只要被调用，就已经假设本地消息是完整的了
func GetMessage(ctx *gin.Context) {
	var requestMap modelview.Message2
	ctx.ShouldBind(requestMap)
	src_user := requestMap.SrcUser //src是对面发消息的客户端，而dst肯定是本机，不用管
	key_word := requestMap.KeyWord
	var messages []model.Message
	if src_user == "" && key_word == "" {
		messages, _ = model.GetMessageByOperate("DecryptRequest2Client")
	} else if src_user == "" {
		messages, _ = model.GetMessageByMap(map[string]string{"Operate": "DecryptRequest2Client", "KeyWord": key_word})
	} else if key_word == "" {
		messages, _ = model.GetMessageByMap(map[string]string{"Operate": "DecryptRequest2Client", "SrcUser": src_user})
	} else {
		messages, _ = model.GetMessageByMap(map[string]string{"Operate": "DecryptRequest2Client", "SrcUser": src_user, "KeyWord": key_word})
	}
	result_messages := make([]modelview.Message, len(messages))
	for i, v := range messages {
		result_messages[i] = modelview.Message{Connect_user: v.SrcUser, App: v.KeyWord}
	}
	ctx.JSON(200, result_messages)
}

// 删除消息
func DeleteMessage(ctx *gin.Context) {
	// 解析表单输入
	var requestMap modelview.Message2
	ctx.ShouldBind(&requestMap)
	src_user := requestMap.SrcUser
	key_word := requestMap.KeyWord
	// 删除消息
	model.DeleteMessage(src_user, username, key_word, "DecryptRequest2Client")
	// 回复前端
	ctx.JSON(200, gin.H{"msg": "delete over."})
}

// 同意消息
func AgreeMessage(ctx *gin.Context) {
	// 解析表单输入
	var requestMap modelview.Message2
	ctx.ShouldBind(&requestMap)
	src_user := requestMap.SrcUser
	key_word := requestMap.KeyWord
	// 查询对应消息
	// msg, _ := model.GetMessage(src_user, dst_user, key_word)
	// operate := msg.Operate
	// // 生成消息返回对应用户
	// new_msg := model.Message{SrcUser: username, DstUser: src_user, KeyWord: key_word, Operate: "Agree " + operate, Params: ""}
	// PostMessage(new_msg)
	// // 删除该消息
	// model.DeleteMessage(src_user, dst_user, key_word)

	//接下来的操作应该是：查询对应消息，然后从params里读出相应的二次加密密文，返回进行解密，将解密后的信息返回
	// 查询对应消息
	msg, _ := model.GetMessage(src_user, username, key_word, "DecryptRequest2Client")
	if msg.DstUser == "" {
		//说明没找到，直接报错
		panic("agree没有找到对应信息:src_user:" + src_user + " key_word:" + key_word)
	}
	passwd1 := msg.Params
	//解密
	stringdata, flag := Deal_B2A_message_to_base(passwd1)
	if flag == 0 {
		ctx.JSON(401, gin.H{"msg": "发生错误"})
		return
	}
	// 生成消息返回对应用户
	new_msg := model.Message{SrcUser: username, DstUser: src_user, KeyWord: key_word, Operate: "DecryptMessage2Server", Params: stringdata}
	PostMessage(new_msg)
	//接收服务器返回的信息
	response := <-ResponseChan
	if response == "AddMessageError1" {
		ctx.JSON(401, gin.H{"msg": "服务器未转发，请重新同意"})
		return
	} else if response == "AddMessagePass" {
		// 删除消息
		model.DeleteMessage(src_user, username, key_word, "DecryptRequest2Client")
		ctx.JSON(200, gin.H{"msg": "同意该请求"})
	}

}

// 不同意消息
func DisagreeMessage(ctx *gin.Context) {
	// 解析表单输入
	var requestMap modelview.Message2
	ctx.ShouldBind(&requestMap)
	src_user := requestMap.SrcUser
	key_word := requestMap.KeyWord
	// // 查询对应消息
	// msg, _ := model.GetMessage(src_user, dst_user, key_word)
	// operate := msg.Operate
	// // 生成消息返回对应用户
	// new_msg := model.Message{SrcUser: username, DstUser: src_user, KeyWord: key_word, Operate: "Disagree " + operate, Params: ""}
	// PostMessage(new_msg)
	// // 删除该消息
	// model.DeleteMessage(src_user, dst_user, key_word)

	//接下来的操作应该是：返回不同意的信息
	new_msg := model.Message{SrcUser: username, DstUser: src_user, KeyWord: key_word, Operate: "DecryptRequestDisAgree2Server", Params: ""}
	PostMessage(new_msg)
	//接收服务器返回的信息
	response := <-ResponseChan
	if response == "AddMessageError1" {
		ctx.JSON(401, gin.H{"msg": "服务器未转发，请重新同意"})
		return
	} else if response == "AddMessagePass" {
		// 删除消息
		model.DeleteMessage(src_user, username, key_word, "DecryptRequest2Client")
		ctx.JSON(200, gin.H{"msg": "不同意该请求"})
	}
}

func GetAddMessage(ctx *gin.Context) {
	// 解析表单输入
	var requestMap modelview.Message2
	ctx.ShouldBind(&requestMap)
	src_user := requestMap.SrcUser
	key_word := requestMap.KeyWord
	// 查询对应添加请求
	var messages []model.Message
	if src_user == "" && key_word == "" {
		messages, _ = model.GetMessageByOperate("EncryptAnnocement2Client")
	} else if src_user == "" {
		messages, _ = model.GetMessageByMap(map[string]string{"Operate": "EncryptAnnocement2Client", "KeyWord": key_word})
	} else if key_word == "" {
		messages, _ = model.GetMessageByMap(map[string]string{"Operate": "EncryptAnnocement2Client", "SrcUser": src_user})
	} else {
		messages, _ = model.GetMessageByMap(map[string]string{"Operate": "EncryptAnnocement2Client", "SrcUser": src_user, "KeyWord": key_word})
	}
	result_messages := make([]modelview.Message, len(messages))
	for i, v := range messages {
		result_messages[i] = modelview.Message{Connect_user: v.SrcUser, App: v.KeyWord}
	}
	// 回复前端
	ctx.JSON(200, result_messages)
}

func DeleteAddMessage(ctx *gin.Context) {
	// 解析表单输入
	var requestMap modelview.Message2
	ctx.ShouldBind(&requestMap)
	src_user := requestMap.SrcUser
	key_word := requestMap.KeyWord
	// 删除对应添加请求
	flag := model.DeleteMessage(src_user, username, key_word, "EncryptAnnocement2Client")
	// 回复前端
	if flag == 1 {
		ctx.JSON(200, gin.H{"msg": "delete over."})
	} else {
		ctx.JSON(401, gin.H{"msg": "delete failed."})
	}
}
