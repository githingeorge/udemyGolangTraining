package main

import (
	"fmt"
)

func visit(numbers []int, cb func(int)) {
	for _, n := range numbers {
		cb(n)
	}
}

func filter(numbers []int, greaterThan func(int) bool) []int {
	var xs []int
	for _, number := range numbers {
		if greaterThan(number) {
			xs = append(xs, number)
		}
	}
	return xs
}
func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n * 2)
	})
	x := []int{1, 2, 3, 4}
	filtered := filter(x, func(n int) bool {
		return n > 1
	})
	fmt.Println(x)
	fmt.Println(filtered)
}
