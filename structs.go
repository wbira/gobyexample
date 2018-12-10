package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 22})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 30}
	fmt.Println(s.name)

	sp := &s
	fmt.Println("before change", sp.age)

	sp.age = 40
	fmt.Println("change age:", sp.age)
	fmt.Println("s also changed", s.age)
	s.name = "John"
	fmt.Println("name s after change", s.name)
	fmt.Println("name sp after change", sp.name)
}
