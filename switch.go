package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("second")
	case 3:
		fmt.Println("third")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend :)")
	default:
		fmt.Println("It's workday :(")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatIAm := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Boolean")
		case int:
			fmt.Println("Integer")
		default:
			fmt.Printf("I don't know type %T\n", t)
		}
	}

	whatIAm(true)
	whatIAm(42)
	whatIAm("Hey Joe")
}
