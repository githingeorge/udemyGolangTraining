package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Print("Enter a number: ")
	fmt.Scan(&a)

	var b int
	fmt.Print("Enter a number: ")
	fmt.Scan(&b)
	if a > b {
		fmt.Printf("%d is greater than %d", a, b)

	} else if a == b {
		fmt.Printf("%d is equal to %d", a, b)

	} else {
		fmt.Printf("%d is less than %d", a, b)

	}
}
