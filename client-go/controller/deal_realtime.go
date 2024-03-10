package controller

import (
	"client-go/model"
)

//该文件用于处理实时的服务器发送过来的信息，根据不同的信息类型，修改不同的数据库

/*
服务器发送过来的消息有以下几种：
1、登录/注册时的用户名和密码，此时它应该可以实时回复，不需要在这里处理
2、回复的公钥（本机会向服务器请求公钥），此时它应该可以实时回复，不需要在这里处理
3、转发解密请求（客户端1向服务器发送解密请求，服务器将该请求转发给本机），Operate=DecryptRequest2Client
4、转发加密通知（客户端1通过本机的公钥加密，则会将该信息发送给本机），Operate=EncryptAnnocement2Client
5、初步解密密文（本机向客户端2发送解密请求，客户端通过服务器将初步解密后的密文发过来），Operate=DecryptMessage2Client
*/

func deal_messages(msg model.Message) int {
	//该函数用于判断不同类型的消息，然后将其转给对应类型的处理函数，返回消息处理的结果，1代表正确
	var flag int
	switch msg.Operate {
	case "DecryptRequest2Client":
		flag = deal_message_DecryptRequest(msg)
	case "EncryptAnnocement2Client":
		flag = deal_message_EncryptAnnocement(msg)
	case "DecryptMessage2Client":
		flag = deal_message_DecryptMessage(msg)
	default:
		panic("接收到的信息类型无法确认")
	}
	if flag == 1 {
		return 1
	} else {
		return 0
	}
}
func deal_message_DecryptRequest(msg model.Message) int {
	//该函数用于处理DecryptRequest请求，具体操作就是将该信息加入message数据库
	return model.AddMessage(msg)
}

func deal_message_EncryptAnnocement(msg model.Message) int {
	//该函数用于处理EncryptAnnocement请求，具体操作就是将该信息加入message数据库
	return model.AddMessage(msg)
}

func deal_message_DecryptMessage(msg model.Message) int {
	//该函数用于处理DecryptMessage请求，具体操作就是修改查询结果表中的信息，将状态改为已完成，然后将密码修改
	return model.ChangeResearchAns(msg.SrcUser, msg.DstUser, msg.KeyWord, "has complete", msg.Params)
}
