package main

import (
	"fmt"
)

func vals() (int, int) {
	return 2, 3
}

func main() {
	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)

}
