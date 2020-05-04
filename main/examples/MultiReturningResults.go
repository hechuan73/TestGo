package main

import "fmt"

func main() {
	a, b := multiResults(1, 2)
	fmt.Println(a, b)
	c, _ := multiResults(1, 2)
	fmt.Println(c)
}

func multiResults(var1 int, var2 int) (int, int) {
	return var1, var2
}