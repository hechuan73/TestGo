package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second*2)
		c1 <- "result 1"
	}()

	// 由于 select 默认处理第一个已准备好的接收操作，如果这个操作超过了允许的 1 秒的话，将会执行超时 case。
	select {
	case msg := <- c1: fmt.Println(msg)
	case <-time.After(time.Second * 1): fmt.Println("timeout 1")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second*1)
		c2 <- "result 2"
	}()

	select {
	case msg := <- c2: fmt.Println(msg)
	case <-time.After(time.Second * 3): fmt.Println("timeout 2")
	}
}
