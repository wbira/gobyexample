package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declared", b)

	var twoDim [2][3]int

	for i := 0; i < 2; i++ {
		for k := 0; k < 3; k++ {
			twoDim[i][k] = i + k
		}
	}

	fmt.Println("2d: ", twoDim)
}
