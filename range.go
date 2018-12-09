package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	m := map[string]string{"a": "apple", "b": "bannan"}
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for k, v := range "go" {
		fmt.Println(k, v)
	}
}
