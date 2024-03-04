// SSE简单实现
// https://blog.csdn.net/Jay_Josby/article/details/128816367
package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var channelMap map[string]chan string

func InitSSE() {
	channelMap = make(map[string]chan string)
}

func AddChannel(key string) chan string {
	// 建立通道
	var channel = make(chan string)
	channelMap[key] = channel
	return channel
}

func BuildNotificationChannel(key string, ctx *gin.Context) {
	// 建立通知频道，对指定key发送实时通知
	channel := AddChannel(key)
	// 更改请求相关参数，更改请求conntent-type类型为流数据
	ctx.Writer.Header().Set("Content-Type", "text/event-stream")
	ctx.Writer.Header().Set("Cache-Control", "no-cache")
	ctx.Writer.Header().Set("Connection", "keep-alive")
	//
	w := ctx.Writer
	flusher, _ := w.(http.Flusher)
	closeNotify := ctx.Request.Context().Done()
	// 创建关闭连接函数，删除通道字典中对应通道
	go func() { // go关键字修饰，代表创建一个轻量级线程异步执行函数
		<-closeNotify           // 等待该通道数据传来，代码阻塞在这里
		delete(channelMap, key) //一旦通道closeNotify有数据，说明请求已结束，连接需关闭，则删除现有信息
		fmt.Println("SSE close for user " + key)
		return
	}() // 紧跟着（）代表立即执行该函数

	fmt.Println("SSE ready for user " + key)
	fmt.Fprintf(w, "data: %s\n\n", "SSE ready for user "+key)
	flusher.Flush()
	// 注意，不能让通道处理跑完，这样消息就终止了，连接就关闭了，后端就无法再通过该连接向前端传信息了
	for msg := range channel { // 从通道中接受数据，直至通道关闭结束循环
		fmt.Fprintf(w, "data:%s\n\n", msg)
		flusher.Flush()
	}

}

// 后端推送消息
func SendNotification(key string, msg string) {
	fmt.Println("SSE send " + msg + " for user " + key)
	channel := channelMap[key]
	msgByte, _ := json.Marshal(msg)
	channel <- string(msgByte)
}
