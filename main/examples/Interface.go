package main

import "fmt"
import "math"

// 定义一个接口
type geometry interface {
	area() float64
	perimeter() float64
}

// 让结构体实现接口
type rect1 struct {
	width, height float64
}

type circle1 struct {
	radius float64
}

func (r rect1) area() float64 {
	return r.width * r.height
}

func (r rect1) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle1) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle1) perimeter() float64 {
	return 2 * math.Pi + c.radius
}

func measure(geo geometry)  {
	fmt.Println(geo)
	fmt.Println(geo.area())
	fmt.Println(geo.perimeter())
}

func main() {
	rec := rect1{10, 20}
	circle := circle1{10}
	measure(rec)
	measure(circle)
}
