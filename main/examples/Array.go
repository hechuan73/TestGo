package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])
	fmt.Println(len(a))

	b := [5]int {1, 2, 3, 4, 5}
	fmt.Println(b)

	var c [2][3]int
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[0]); j++ {
			c[i][j] = i+j
		}
	}

	fmt.Println(c)
}
