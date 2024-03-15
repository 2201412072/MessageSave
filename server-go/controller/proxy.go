package controller

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"server-go/model"
)

var readers = make(map[string]*bufio.Reader)
var writers = make(map[string]*bufio.Writer)
var connmap = make(map[string]*net.Conn) //该对象用于记录每个用户的连接，用于后续的连接释放

func InitProxy() {
	fmt.Println("InitProxy:", server_addr+":"+port)
	// 监听新连接
	listener, err := net.Listen("tcp", server_addr+":"+port)
	fmt.Println(listener)
	if err != nil {
		fmt.Println("Listen Error! ", err)
	}
	defer listener.Close()
	// 不断接受客户端连接
	for {
		ConnAccept(listener)
	}
}

// 接受客户端连接
func ConnAccept(listener net.Listener) {
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Failed to create connection. ", err)
	}
	//不能延时释放它，因为运行到函数最后，会开一个新协程，此时就会执行defer，就把conn删掉了，所以最好能够在客户端断开连接以后再close
	//但是单纯先记录conn，再删掉conn不行，因为删conn时会将该tcp连接断开，所以应该是指针实现
	//defer conn.Close()
	// 获取客户端用户名
	reader := bufio.NewReader(conn)
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Socket recive error ", err)
	}

	var userinfo = make(map[string]string)
	json.Unmarshal(buffer[:n], &userinfo)
	fmt.Println(userinfo)
	user := userinfo["user"]
	//标志着该用户在线
	model.UpdateStageUserPassword(user, 1)
	connmap[user] = &conn
	// 准备接收消息
	writer := bufio.NewWriter(conn)
	readers[user] = reader
	writers[user] = writer
	// 另起一个线程接收消息
	go func() {
		for {
			message, _ := reader.ReadString('\n')
			if len(message) == 0 {
				//可能是tcp相关控制的包，目前发现的情况有，tcp链路被断开，则message为空
				ConnDelete(user)
				//标志着该用户离线
				model.UpdateStageUserPassword(user, 0)
				fmt.Println("用户 ", user, " 中断了通话")
				break
			}
			fmt.Printf("Get %v len %v message: %v \n", user, len(message), message)
			// 处理接收到的消息
			var msg model.Message
			json.Unmarshal([]byte(message), &msg)
			deal_messages(msg)
		}
	}()
}

// 发送消息至关联用户
func PostMessage(message model.Message) int {
	if message.SrcUser == "" {
		message.SrcUser = "server"
	}
	meesageData, _ := json.Marshal(message)
	ConnSend(message.DstUser, string(meesageData))
	return 1
}

// 发送消息
func ConnSend(user string, msg string) error {
	fmt.Println("向", user, "发送消息:", msg)
	writer, exist := writers[user]
	if exist {
		writer.Write([]byte(msg))
		writer.Flush()
		return nil
	} else {
		return errors.New("login user isn't exist")
	}
}

// 删除该连接
func ConnDelete(user string) {
	(*connmap[user]).Close()
	delete(connmap, user)
}
