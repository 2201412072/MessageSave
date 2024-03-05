package controller

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

var readers map[string]*bufio.Reader
var writers map[string]*bufio.Writer

func InitProxy() {
	// 监听新连接
	listener, err := net.Listen("tcp", server_addr+":"+port)
	if err != nil {
		fmt.Println("Listen Error! ", err)
	}
	defer listener.Close()
	// 不断接受客户端连接
	go func() {
		for {
			ConnAccept(listener)
		}
	}()
}

// 接受客户端连接
func ConnAccept(listener net.Listener) {
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Failed to create connection. ", err)
	}
	defer conn.Close()
	// 获取客户端用户名
	reader := bufio.NewReader(conn)
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Socket recive error ", err)
	}
	var userinfo map[string]string
	json.Unmarshal(buffer[:n], &userinfo)
	user := userinfo["user"]
	// 准备接收消息
	writer := bufio.NewWriter(conn)
	readers[user] = reader
	writers[user] = writer
	// 另起一个线程接收消息
	go func() {
		for {
			message, _ := reader.ReadString('\n')
			fmt.Println("Get %s message: %s", user, message)
			// 处理接收到的消息
		}
	}()
}

// 发送消息
func ConnSend(user string, msg string) {
	writer := writers[user]
	writer.Write([]byte(msg))
	writer.Flush()
}
