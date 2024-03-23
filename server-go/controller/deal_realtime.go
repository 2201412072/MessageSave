package controller

import (
	"fmt"
	"server-go/model"
	"server-go/util"
)

//该文件用于处理实时的客户端送过来的信息，根据不同的信息类型，修改不同的数据库

/*
客户端发送过来的消息有以下几种：
1、发送过来的登录的信息，Operate=Login2Server
2、发送过来的注册的信息，Operate=Register2Server
3、发送过来的请求公钥的信息，Operate=PublicKeyRequest2Server
4、发送过来的解密请求信息（客户端1希望向客户端2发送解密请求，在服务器端中转）,Operate=DecryptRequest2Server
5、发送过来的加密通知（客户端1通过客户端2公钥加密，向客户端2发送该通知，在服务器端中转），Operate=EncryptAnnocement2Server
6、回复的初步解密密文（客户端1让客户端2解密，客户端2发送回初步解密密文，在服务器端中转），Operate=DecryptMessage2Server
7、返回的不同意解密的信息（客户端1希望向客户端2发送解密请求，客户端2不同意，将该信息返回给客户端1，在服务器端中转），Operate=DecryptRequestDisAgree2Server
8、发送过来的公钥信息，服务器需要记录一下，Operate=PublicKeyRecord2Server
*/

func deal_messages(msg model.Message) int {
	//该函数用于判断不同类型的消息，然后将其转给对应类型的处理函数，返回消息处理的结果，1代表正确
	fmt.Println("deal_messages")
	var flag int
	switch msg.Operate {
	case "Login2Server":
		flag = deal_message_Login(msg)
	case "Register2Server":
		flag = deal_message_Register(msg)
	case "PublicKeyRequest2Server":
		flag = deal_message_PublicKeyRequest(msg)
	case "DecryptRequest2Server":
		flag = deal_message_DecryptRequest(msg)
	case "EncryptAnnocement2Server":
		flag = deal_message_EncryptAnnocement(msg)
	case "DecryptMessage2Server":
		flag = deal_message_DecryptMessage(msg)
	case "PublicKeyRecord2Server":
		flag = deal_message_PublicKeyRecord(msg)
	case "DecryptRequestDisAgree2Server":
		flag = deal_message_DecryptRequestDisAgree(msg)
	default:
		panic("接收到的信息类型无法确认")
	}
	return flag
}

func deal_message_Login(msg model.Message) int {
	//该函数处理Login2Server请求，具体操为查看用户表中是否有该用户，同时密码是否正确，然后将结果发送给客户端
	//如果正确返回1
	temp, err := model.GetUserPassword(msg.SrcUser)
	if err == 0 {
		//没找到该用户
		ConnSend(msg.SrcUser, "LoginError1") //can't find this user name
		return 0
	}
	if temp.Password != msg.KeyWord {
		//密码不匹配
		ConnSend(msg.SrcUser, "LoginError2") //password inputted is wrong
		return 0
	}
	ConnSend(msg.SrcUser, "LoginPass")
	return 1
}

func deal_message_Register(msg model.Message) int {
	//该函数处理Register2Server请求，具体操作为查看用户表中是否有该用户，如果没有就加进去，然后将结果发送给客户端
	_, err := model.GetUserPassword(msg.SrcUser)
	if err == 1 {
		//已经有一个用户了
		ConnSend(msg.SrcUser, "RegisterError1") //already have this username
		return 0
	}
	err = model.AddUserPassword(msg.SrcUser, msg.KeyWord)
	if err != 1 {
		//添加用户失败
		ConnSend(msg.SrcUser, "RegisterError1") //already have this username
		return 0
	}
	ConnSend(msg.SrcUser, "RegisterPass")
	return 1
}

func deal_message_PublicKeyRequest(msg model.Message) int {
	//该函数处理PublicKeyRequest2Server请求，具体操作为查看用户表中是否有该用户，将其公钥发送回客户端
	temp, err := model.GetPublicKeyByUser(msg.SrcUser)
	if err != 1 {
		//没找到用户
		ConnSend(msg.SrcUser, "PublicKeyRequestError1")
		return 0
	}
	ans, _ := util.Bytes2base_string(temp)
	ConnSend(msg.SrcUser, "PublicKeyRequestPass")
	ConnSend(msg.SrcUser, ans)
	return 1
}

func deal_message_DecryptRequest(msg model.Message) int {
	//该函数处理DecryptRequest2Server，具体操作为加入暂存消息表中,如果目标用户连着服务器，则顺便转发给它
	err := model.AddMessage(msg)
	if err != 1 {
		//加入出错
		ConnSend(msg.SrcUser, "AddMessageError1")
		return 0
	}
	// 向源用户发送“添加成功”
	src_user := msg.SrcUser
	ConnSend(src_user, "AddMessagePass")
	// 向目的用户发送添加密码请求
	msg.Operate = "DecryptRequest2Client"
	msg.SrcUser = msg.DstUser
	msg.DstUser = src_user
	if send_realtime_messge(msg) == 1 {
		//在线，发送过去了
		model.DeleteMessage(msg.SrcUser, msg.DstUser, msg.KeyWord, msg.Operate)
	}
	return 1
}

func deal_message_EncryptAnnocement(msg model.Message) int {
	//该函数处理EncryptAnnocement2Server，具体操作为加入暂存消息表中,如果目标用户连着服务器，则顺便转发给它
	err := model.AddMessage(msg)
	if err != 1 {
		//加入出错
		ConnSend(msg.SrcUser, "AddMessageError1")
		return 0
	}
	ConnSend(msg.SrcUser, "AddMessagePass")
	msg.Operate = "EncryptAnnocement2Client"
	if send_realtime_messge(msg) == 1 {
		//在线，发送过去了
		model.DeleteMessage(msg.SrcUser, msg.DstUser, msg.KeyWord, msg.Operate)
	}
	return 1
}

func deal_message_DecryptMessage(msg model.Message) int {
	//该函数处理DecryptMessage2Server，具体操作为加入暂存消息表中,如果目标用户连着服务器，则顺便转发给它
	err := model.AddMessage(msg)
	if err != 1 {
		//加入出错
		ConnSend(msg.SrcUser, "AddMessageError1")
		return 0
	}
	ConnSend(msg.SrcUser, "AddMessagePass")
	msg.Operate = "DecryptMessage2Client"
	if send_realtime_messge(msg) == 1 {
		//在线，发送过去了
		model.DeleteMessage(msg.SrcUser, msg.DstUser, msg.KeyWord, msg.Operate)
	}
	return 1
}

func deal_message_PublicKeyRecord(msg model.Message) int {
	//该函数处理deal_message_PublicKeyRecord2Server，具体操作为将msg中的公钥放入公钥数据库中
	_, err := model.GetPublicKeyByUser(msg.SrcUser)
	if err == 1 {
		//公钥表中已经存在该用户的公钥
		ConnSend(msg.SrcUser, "PublicKeyRecordError1")
		return 0
	}
	temp_public_key, _ := util.Base_string2bytes(msg.Params)
	model.AddPublicKey(msg.SrcUser, temp_public_key)
	ConnSend(msg.SrcUser, "PublicKeyRecordPass")
	return 1
}

// 该函数用于判断服务器是否连接着目标用户，如果连接就将该消息发送过去。假如没有发送成功呢？
func send_realtime_messge(msg model.Message) int {
	temp, _ := model.GetUserPassword(msg.DstUser)
	if temp.Stage == 0 {
		//离线
		return 0
	} else {
		//在线
		return PostMessage(msg)
	}
}

func deal_message_DecryptRequestDisAgree(msg model.Message) int {
	//该函数用来处理DecryptRequestDisAgree2Server，具体操作为加入暂存消息表中,如果目标用户连着服务器，则顺便转发给它
	err := model.AddMessage(msg)
	if err != 1 {
		//加入出错
		ConnSend(msg.SrcUser, "AddMessageError1")
		return 0
	}
	ConnSend(msg.SrcUser, "AddMessagePass")
	msg.Operate = "DecryptRequestDisAgree2Client"
	if send_realtime_messge(msg) == 1 {
		//在线，发送过去了
		model.DeleteMessage(msg.SrcUser, msg.DstUser, msg.KeyWord, msg.Operate)
	}
	return 1
}
