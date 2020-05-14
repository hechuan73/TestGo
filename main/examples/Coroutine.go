package main

import "fmt"

func main() {

	// 直接调用
	f("direct")

	// 启动协程
	go f("goroutine")

	// 协程执行匿名函数
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}