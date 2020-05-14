package main

import "fmt"

func main() {
	channel := make(chan string, 3)
	channel <- "1"
	channel <- "2"
	channel <- "3"
	// 带缓冲的通道，可以先关闭再遍历
	close(channel)

	for elem := range channel {
		fmt.Print(elem)
	}
}
