package controller

import (
	"bufio"
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
	// 接受客户端连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Failed to create connection. ", err)
	}
	defer conn.Close()
	// 读取发送消息
	/*
		go func() {
			// 创建读写缓冲区
			reader := bufio.NewReader(conn)
			writer := bufio.NewWriter(conn)
			// readers["new user"]=reader
			// writers["1"]=writer
			// 读取消息
			go func() {

			}()
			// 发送消息
			for {
				writer.WriteString("new message")
				writer.Flush()
			}
		}()
	*/
}
