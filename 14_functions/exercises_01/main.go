package main

import (
	"fmt"
)

func main() {
	isEven := func(number int) (int, bool) {

		return number / 2, number%2 == 0
	}

	number, even := isEven(5)
	fmt.Println(number, even)

	number, even = isEven(4)
	fmt.Println(number, even)
	greatest(1, 3, 5, 43, 41, 58, 21, 32)
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 3, 54, 5}
	foo(aSlice...)
	foo()
}

func greatest(numbers ...int) {
	var g int
	for _, n := range numbers {
		if n > g {
			g = n
		}
	}
	fmt.Println(g)
}

func foo(numbers ...int) {
	fmt.Println(numbers)
}
