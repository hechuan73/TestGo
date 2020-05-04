package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _,num := range nums {
		sum += num
	}
	fmt.Println("Sum: ", sum)

	for i := range nums {
		fmt.Println("Index: ", i)
	}

	m := map[string]int{"k1":1, "k2":2}
	for k, v := range m {
		fmt.Printf("%s -> %d\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}


}
