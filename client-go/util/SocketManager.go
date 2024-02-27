package util

import (
	"fmt"
	"net"
)

var conn *net.TCPConn

// 新建socket连接
func NewConn(server_addr string, port string) *net.TCPConn {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server_addr)
	if err != nil {
		fmt.Println("Connection create error! ", err)
		return nil
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	return conn
}

// 发送消息
func ConnSend(msg string) {
	conn.Write([]byte(msg))
}

// 接受消息
func ConnRecive() string {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Socket recive error ", err)
	}
	msg := string(buffer[:n])
	return msg
}

// 关闭连接
func ConnClose() {
	conn.Close()
}
