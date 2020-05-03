package main

import "fmt"

func main() {
	if 7&1 ==0 {
		fmt.Println("even")
	} else if 7&1 !=0 {
		fmt.Println("odd")
	}
}
