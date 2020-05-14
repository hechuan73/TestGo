package main

import "fmt"

func main() {
	message := make(chan string)
	signal := make(chan string)

	// select加上default分支后，可以不用阻塞等待任意一个case分支完成，而直接走default分支，实现非阻塞

	// message没有数据，走default分支
	select {
	case msg := <- message:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}


	// message没有缓冲，所以发送操作没法处理，需要有接收端准备好才能发送，所以这里仍然走default分支。如果加上缓冲，则可以进行发送
	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <- message:
		fmt.Println("received message", msg)
	case sig := <- signal:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
