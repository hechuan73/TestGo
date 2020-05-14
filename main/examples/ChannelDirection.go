package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
	// 从只发通道中接收数据会编译报错
	// str := <- pings
}

// 形参的顺序不能写反了，因为程序中是先接收值，所以接收通道要写在前面，不然会报死锁错误
func pong(pings <-chan string, pongs chan<- string) {
	msg := <- pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}
