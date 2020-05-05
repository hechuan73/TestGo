package main

import "fmt"

func main() {
	fmt.Println(person{"bob", 10})
	fmt.Println(person{name: "tony", age: 20})
	fmt.Println(person{name: "terry"})

	fmt.Println(&person{"judy", 10})

	s := person{name: "60"}
	s.name = "lily"
	s.age = 20
	fmt.Println(s)

	l := s
	l.name = "king"
	fmt.Println(l)

}

type person struct {
	name string
	age int
}
