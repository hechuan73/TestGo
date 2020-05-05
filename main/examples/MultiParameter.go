package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3))
	nums := []int {1, 2, 3}
	// slice 要加"..."调用
	fmt.Println(sum(nums...))
}

func sum(nums ...int) int {
	res := 0
	for _, num := range nums{
		res += num
	}

	return res
}