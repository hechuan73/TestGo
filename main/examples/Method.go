package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return r.height * 2 + r.width * 2
}

func main() {
	rect := rect{5, 10}
	fmt.Println(rect.area())
	fmt.Println(rect.perimeter())

	c := rect
	fmt.Println(c.area())
	fmt.Println(c.perimeter())
}
