package main

import "fmt"
import "time"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second*1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second*2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// 通道选择，类似于多路复用
		select {
		case msg := <-c1:
			fmt.Println(msg)
		case msg := <-c2:
			fmt.Println(msg)
		}
	}
}
