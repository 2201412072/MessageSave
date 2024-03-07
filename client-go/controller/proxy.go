package controller

import (
	"client-go/model"
	"client-go/util"
	"container/list"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/rpc"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/antage/eventsource.v1"
)

var client *rpc.Client
var msgChannel chan string
var msgQueue list.List
var es eventsource.EventSource
var err error

func InitProxy() {
	// 初始化消息代理
	// 创建和服务器的连接
	// client, err = rpc.DialHTTP("tcp", server_addr+port)
	// if err != nil {
	// 	fmt.Println("Error! ", err)
	// }
	util.NewConn(server_addr, port)
	// 创建SSE连接，从服务器获取消息并传给前端
	es = util.NewEventSource("proxy_message", "/ProxyMessage/events", RecvMessage)
	// // 获取服务器端积压的数据
	// messages, flag := GetMessage()
}

// func GetMessage() ([]model.Message, int) { // RPC远程调用实现获取消息功能
// 	// 启动代理时从服务器端拉取消息
// 	// 远程调用服务器获取消息函数
// 	var messages []model.Message
// 	err = client.Call("Server.GetMessageByUser", username, &messages) // 同步调用服务器GetMessageByUser函数（查询完就删掉对应表项）
// 	if err != nil {
// 		fmt.Println("Error! ", err)
// 	}
// 	return messages, 1
// }

// 获取未处理的消息
func GetMessage(ctx *gin.Context) {
	// 获取本地缓存的未处理的消息
	local_messages, _ := model.GetMessages()
	// 获取服务器消息并缓存
	var net_messages []model.Message
	util.ConnSend("GetMessageByUser " + username)
	msg_json := util.ConnRecive()
	err := json.Unmarshal([]byte(msg_json), &net_messages)
	if err != nil {
		fmt.Println("Error! ", err)
	}
	for _, msg := range net_messages { // 缓存
		model.AddMessage(msg)
	}
	// 合并两堆消息
	messages := append(local_messages, net_messages...)
	// 消息队列
	// for _,msg := range messages {
	// 	msgQueue.PushBack(msg)
	// }
	// 回复前端
	ctx.JSON(200, messages)
}

// 删除消息
func DeleteMessage(ctx *gin.Context) {
	// 解析表单输入
	var requestMap model.Message
	ctx.ShouldBind(&requestMap)
	src_user := requestMap.SrcUser
	dst_user := requestMap.DstUser
	key_word := requestMap.KeyWord
	// 删除消息
	model.DeleteMessage(src_user, dst_user, key_word)
	// 回复前端
	ctx.JSON(200, gin.H{"msg": "delete over."})
}

// 同意消息
func AgreeMessage(ctx *gin.Context) {
	// 解析表单输入
	var requestMap model.Message
	src_user := requestMap.SrcUser
	dst_user := requestMap.DstUser
	key_word := requestMap.KeyWord
	// 查询对应消息
	msg, _ := model.GetMessage(src_user, dst_user, key_word)
	operate := msg.Operate
	// 生成消息返回对应用户
	new_msg := model.Message{SrcUser: username, DstUser: src_user, KeyWord: key_word, Operate: "Agree " + operate, Params: ""}
	PostMessage(new_msg)
	// 删除该消息
	model.DeleteMessage(src_user, dst_user, key_word)
}

// 不同意消息
func DisagreeMessage(ctx *gin.Context) {
	// 解析表单输入
	var requestMap model.Message
	src_user := requestMap.SrcUser
	dst_user := requestMap.DstUser
	key_word := requestMap.KeyWord
	// 查询对应消息
	msg, _ := model.GetMessage(src_user, dst_user, key_word)
	operate := msg.Operate
	// 生成消息返回对应用户
	new_msg := model.Message{SrcUser: username, DstUser: src_user, KeyWord: key_word, Operate: "Disagree " + operate, Params: ""}
	PostMessage(new_msg)
	// 删除该消息
	model.DeleteMessage(src_user, dst_user, key_word)
}

func RecvMessage() {
	// 代理启动后接收服务器传来的消息（不知道怎么实现才好，就先还是查询message表吧）
	// var messages []model.Message
	// err = client.Call("Server.GetMessageByUser", username, &messages) // 同步调用服务器GetMessageByUser函数
	// if err != nil {
	// 	fmt.Println("Error! ", err)
	// }

	for {
		var messages []model.Message
		msg_json := util.ConnRecive()
		err := json.Unmarshal([]byte(msg_json), &messages)
		if err != nil {
			fmt.Println("Error! ", err)
			return
		}
		for msg := range messages {
			msgQueue.PushBack(msg)
		}
		//这里是处理message的函数，要求能够针对不同类型的message来修改不同数据库
		for msg := range message {
			deal_messages(msg)
		}
		// 后端通知前端msgQueue更新
		// 只设置发送数据，不添加事件名
		es.SendEventMessage(fmt.Sprintf("send data: %s", time.Now().Format("2006-01-02 15:04:05")), "time", "")
		es.SendEventMessage(string(msgQueue.Len()), "msgQueue len", "")
	}
	// 后端主动向前端推送消息，使用SSE（不用WebSocket了）
	// for msg := range messages {
	// 	// msgByte, _ := json.Marshal(msg)
	// 	// msgChannel <- string(msgByte)

	// 	// 只设置发送数据，不添加事件名
	// 	es.SendEventMessage(fmt.Sprintf("send data: %s", time.Now().Format("2006-01-02 15:04:05")), "time", "")

	// 	// es.SendEventMessage(msg, "msg", "")
	// }
}

// 发送消息中转服务器至关联用户
func PostMessage(message model.Message) int {
	meesageData, _ := json.Marshal(message)
	message_json := string(meesageData)
	util.ConnSend(message_json)
	return 1
}

// 查询自身数据库有无对应公钥
func GetPublicKeyByUser(user string) (*rsa.PublicKey, int) {
	public_key, _ := model.GetPublicKeyByUser(user)
	if public_key == nil {
		// 假如没有，向服务器请求
		message := model.Message{SrcUser: username, DstUser: user, KeyWord: "", Operate: "Search PublicKey", Params: ""}
		PostMessage(message)
		recv_json := util.ConnRecive()
		var recv_msg model.Message
		json.Unmarshal([]byte(recv_json), &recv_msg)
		public_key = []byte(recv_msg.Params)
		// 存储至数据库
		model.AddPublicKey(user, public_key)
	}
	public_key_ans, _ := util.PublicKey_from_bytes(public_key)
	return public_key_ans, 1
}

// 请求指定用户解码
func RequireDecrypt(user string, application string, passwd2 string) int {
	message := model.Message{SrcUser: username, DstUser: user, KeyWord: application, Operate: "Decrypt", Params: ""}
	PostMessage(message)
	return 1
}
