package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a";s[1] = "b";s[2] = "c"
	fmt.Println(s)
	fmt.Println(s[2])

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("s: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c: ", c)

	l := s[2:5]
	fmt.Println("l1: ", l)

	l = s[:3]
	fmt.Println("l2: ", l)

	l = s[2:]
	fmt.Println("l3: ", l)

	t := []string {"g", "h", "i"}
	fmt.Println("t: ", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i+1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2D: ", twoD)
}
