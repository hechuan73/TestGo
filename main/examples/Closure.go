package main

import "fmt"

// 注意：闭包这里要写两个"func"
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}


