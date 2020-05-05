package main

import "fmt"

func main() {
	fmt.Println(cal(7))
}

func cal(a int) int {
	if a == 0 {
		return 1
	}

	return a * cal(a-1)
}