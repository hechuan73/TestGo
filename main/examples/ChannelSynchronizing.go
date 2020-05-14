package main

import "fmt"
import "time"

func work(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func test() {
	fmt.Println("test")
}
func main() {
	go test()
	done := make(chan bool, 1)
	go work(done)
	// 阻塞等待。如果不等待，主协程将立即退出，看不到另一个协程打印的结果
	// 但在Java中，主线程退出，其它线程会打印出结果
	<- done
}
