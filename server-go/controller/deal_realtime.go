package controller

import (
	"server-go/model"
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
*/

func deal_messages(msg model.Message) int {
	//该函数用于判断不同类型的消息，然后将其转给对应类型的处理函数，返回消息处理的结果，1代表正确
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
	default:
		panic("接收到的信息类型无法确认")
	}
}

func deal_message_Login(msg model.Message) int {
	//该函数处理Login2Server请求，具体操为查看用户表中是否有该用户，同时密码是否正确，如果正确返回1
	return 1
}

func deal_message_Register(msg model.Message) int {
	//该函数处理Register2Server请求，具体操作为查看用户表中是否有该用户，如果没有就加进去
	return 1
}

func deal_message_PublicKeyRequest(msg model.Message) int {
	//该函数处理PublicKeyRequest2Server请求，具体操作为查看用户表中是否有该用户，将其公钥发送回客户端
	return 1
}

func deal_message_DecryptRequest(msg model.Message) int {
	//该函数处理DecryptRequest2Server，具体操作为加入暂存消息表中,如果目标用户连着服务器，则顺便转发给它
}

func deal_message_EncryptAnnocement(msg model.Message) int {
	//该函数处理EncryptAnnocement2Server，具体操作为加入暂存消息表中,如果目标用户连着服务器，则顺便转发给它
}

func deal_message_DecryptMessage(msg model.Message) int {
	//该函数处理DecryptMessage2Server，具体操作为加入暂存消息表中,如果目标用户连着服务器，则顺便转发给它
}
