package main

import "fmt"

func main() {

	// 设置带缓存的通道，缓存大小为"值的个数"，以下例子表示可以缓存两个值
	channel := make(chan string, 2)
	// 缓存值
	channel <- "message1"
	channel <- "message2"

	fmt.Println(<- channel)
	fmt.Println(<- channel)
	// 如果接收端读取次数超过缓存数，会造成死锁等待
	// fmt.Println(<- channel)
}
