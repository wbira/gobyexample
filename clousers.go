package main

import (
	"fmt"
)

func initSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := initSequence()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	oneMoreTime := initSequence()
	fmt.Println(oneMoreTime())

}
