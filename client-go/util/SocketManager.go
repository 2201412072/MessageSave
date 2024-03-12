package util

import (
	"fmt"
	"net"
)

var conn *net.TCPConn
var buffer = make([]byte, 1024)

// 新建socket连接
func NewConn(server_addr string, port string) *net.TCPConn {
	tcpAddr, err := net.ResolveTCPAddr("tcp", server_addr+":"+port)
	if err != nil {
		fmt.Println("Connection create error! ", err)
		return nil
	}
	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	fmt.Println("conn:", conn)
	return conn
}

// 发送消息
func ConnSend(msg string) {
	fmt.Println("向服务器发送消息:", msg)
	conn.Write([]byte(msg))
}

// 接受消息
func ConnRecive() string {
	//buffer := make([]byte, 1024)
	fmt.Println("conn:", conn)
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
