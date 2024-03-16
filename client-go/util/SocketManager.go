package util

import (
	"bufio"
	"fmt"
	"net"
)

var conn *net.TCPConn
var writer *bufio.Writer
var reader *bufio.Reader
var buffer = make([]byte, 1024)

// 新建socket连接
func NewConn(server_addr string, port string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", server_addr+":"+port)
	if err != nil {
		fmt.Println("Connection create error! ", err)
		return
	}
	conn, err = net.DialTCP("tcp", nil, tcpAddr)
	fmt.Println("conn:", conn)
	writer = bufio.NewWriter(conn)
	reader = bufio.NewReader(conn)
}

// 发送消息
func ConnSend(msg string) {
	fmt.Println("向服务器发送消息:", msg)
	// conn.Write([]byte(msg))
	writer.Write([]byte(msg + "/n"))
	writer.Flush()
}

// 接受消息
func ConnRecive() (string, error) {
	// buffer := make([]byte, 1024)
	/*
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Socket recive error ", err)
			return "", err
		}
		msg := string(buffer[:n])
	*/
	msg, _ := reader.ReadString('\n')
	fmt.Println("Conn recevice ", msg)
	return msg, nil
}

// 关闭连接
func ConnClose() {
	conn.Close()
}
