package main

import "fmt"

func main() {
	i := 1
	// 传值
	zeroVal(i)
	fmt.Println(i)

	// 传引用
	zeroPtr(&i)
	fmt.Println(i)

	// 打印引用地址
	fmt.Println(&i)
}

func zeroVal(a int) {
	a = 0
}

func zeroPtr(a *int) {
	*a = 0
}
