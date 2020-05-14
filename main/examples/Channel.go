package main

import "fmt"

func main() {

	// 创建channel，默认发送和接收都是阻塞的，所以在发送和接收时不能并行做其他事情。同时必须接收端准备好后才能开始发送。
	channel := make(chan string)
	// 使用协程发送数据
	go func() {channel <- "message"}()
	// 接收数据
	msg := <- channel
	fmt.Println(msg)

}
