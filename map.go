package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	fmt.Println("m:", m)

	val := m["a"]
	fmt.Println("get:", val)
	fmt.Println("len:", len(m))

	delete(m, "b")
	fmt.Println(m)

	_, isC := m["c"]
	fmt.Println("isC:", isC)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
