package util

import (
	"fmt"
	"net/http"

	"gopkg.in/antage/eventsource.v1"
)

var eventsourceMap map[string]eventsource.EventSource

func NewEventSource(key string, url string, f func()) eventsource.EventSource {
	es := eventsource.New(nil, nil)
	eventsourceMap[key] = es
	defer es.Close()
	// 对应url处理函数
	http.Handle(url, es)
	// 执行回调函数
	go f()
	fmt.Println("Open url " + url + "for user " + key)
	return es
}

func GetEventSource(key string) eventsource.EventSource {
	return eventsourceMap[key]
}
