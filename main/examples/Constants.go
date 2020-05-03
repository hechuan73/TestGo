package main

import "fmt"
import "math"

const s string = "global constant"
func main()  {
	fmt.Println(s)
	const str string = "partial constant"
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))


}
