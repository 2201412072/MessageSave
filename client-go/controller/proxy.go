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

	"gopkg.in/antage/eventsource.v1"
)

var client *rpc.Client
var msgChannel chan string
var msgQueue list.List

// var ResponseMsgQueue list.List //这是回应消息队列
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
	//发送用户名
	temp := make(map[string]string)
	temp["user"] = username
	tempbyte, _ := json.Marshal(temp)
	util.ConnSend(string(tempbyte))

	// 创建SSE连接，从服务器获取消息并传给前端
	es = util.NewEventSource("proxy_message", "/ProxyMessage/events", RecvMessage)

	// 发送自己的公钥，感觉没必要，注册的时候发送就行。
	go func() {
		//之所以新开一个线程，是因为recvmessage在initproxy线程，如果不新开，那么没办法从recvmessage那里收到消息
		//发送自身公钥
		var message model.Message
		message.SrcUser = username
		message.DstUser = "server"
		message.Operate = "PublicKeyRecord2Server"
		tempbyte, _ = PublicKeyBytesFromFile()
		message.Params, _ = util.Bytes2base_string(tempbyte)
		PostMessage(message)
		//接收来自服务器的反馈，即服务器是否正确接收到了公钥
		response := <-ResponseChan
		switch response {
		case "PublicKeyRecordError1":
			fmt.Println("服务器已存在该公钥，此次发送无效")
		case "PublicKeyRecordPass":
			fmt.Println("服务器接收成功")
		default:
			fmt.Println("接收到未知信息，", response)
		}
	}()

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
func PullMessage() error {
	// // 获取本地缓存的未处理的消息
	// local_messages, _ := model.GetMessages()
	// 获取服务器消息并缓存
	var net_messages []model.Message
	util.ConnSend("GetMessageByUser " + username)
	msg_json, err := util.ConnRecive()
	if err != nil { // 连接出错
		fmt.Println("Socket Connect Error!")
		return err
	}
	err = json.Unmarshal([]byte(msg_json), &net_messages)
	if err != nil {
		fmt.Println("Error! ", err)
	}
	for _, msg := range net_messages { // 缓存
		model.AddMessage(msg)
	}
	// 合并两堆消息
	// messages := append(local_messages, net_messages...)
	// return messages, nil
	return nil
}

func RecvMessage() {
	// 代理启动后接收服务器传来的消息（不知道怎么实现才好，就先还是查询message表吧）
	// var messages []model.Message
	// err = client.Call("Server.GetMessageByUser", username, &messages) // 同步调用服务器GetMessageByUser函数
	// if err != nil {
	// 	fmt.Println("Error! ", err)
	// }

	// 用户A向用户A发送解密请求，服务器会向用户A发送请求发送成功，再向用户A发送请求解密，但是目前用户A似乎只能接收到前面的消息，后面的接收不到了。
	for {
		var messages []model.Message
		msg_json, err := util.ConnRecive()
		if err != nil {
			fmt.Println("Socket Connect Error!")
			return
		}
		fmt.Println("msg:", msg_json)
		if len(msg_json) == 0 {
			//目前看的是keep alive报文，不用管
			continue
		}
		err = json.Unmarshal([]byte(msg_json), &messages)
		if err != nil {
			//说明是回应消息，目前的处理方式是：将其直接通过通道送入
			ResponseChan <- msg_json
			continue
			// else {
			// 	fmt.Println("Error:", err)
			// 	return
			// }
		}
		for _, msg := range messages {
			msgQueue.PushBack(msg)
		}
		//这里是处理message的函数，要求能够针对不同类型的message来修改不同数据库
		for _, msg := range messages {
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
	if message.SrcUser == "" {
		message.SrcUser = username
	}
	var temp = make([]model.Message, 1)
	temp[0] = message
	meesageData, _ := json.Marshal(temp)
	//该\n是由于服务器接收数据包时，按照\n结尾（目前放在connsend部分）
	message_json := string(meesageData)
	util.ConnSend(message_json)
	return 1
}

// 查询自身数据库有无对应公钥
func GetPublicKeyByUser(user string) (*rsa.PublicKey, int) {
	public_key, _ := model.GetPublicKeyByUser(user)
	if len(public_key) == 0 {
		// 假如没有，向服务器请求
		message := model.Message{SrcUser: username, DstUser: user, KeyWord: "", Operate: "PublicKeyRequest2Server", Params: ""}
		PostMessage(message)
		// recv_json := util.ConnRecive()
		// var recv_msg model.Message
		// json.Unmarshal([]byte(recv_json), &recv_msg)
		// public_key = []byte(recv_msg.Params)
		//此处在发送消息后，需要接收回应消息
		response := <-ResponseChan
		fmt.Println("response:", response)
		if response == "PublicKeyRequestError1" {
			//服务器没有存储该用户的公钥，完犊子了
			return nil, 2
		} else if response == "PublicKeyRequestPass" {
			response := <-ResponseChan //此时收的数据才是真正的公钥
			temp_public_key, _ := util.Base_string2bytes(response)
			// 存储至数据库
			model.AddPublicKey(user, temp_public_key)
			public_key_ans, _ := util.PublicKey_from_bytes(public_key)
			return public_key_ans, 1
		} else {
			return nil, 0
		}
	}
	fmt.Println(public_key)
	public_key_ans, _ := util.PublicKey_from_bytes(public_key)
	return public_key_ans, 1
}

// 请求指定用户解码
func RequireDecrypt(user string, application string, passwd2 string) int {
	message := model.Message{SrcUser: username, DstUser: user, KeyWord: application, Operate: "DecryptRequest2Server", Params: passwd2}
	PostMessage(message)
	return 1
}
