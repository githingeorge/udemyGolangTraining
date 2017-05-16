package main

import (
	"fmt"
)

func main() {
	n := average(1, 2, 3, 6.5)
	fmt.Print(n)

}

func average(numbers ...float64) float64 {
	fmt.Printf("%T - %v \n", numbers, numbers)
	var total float64
	for _, number := range numbers {
		total += number
	}

	fmt.Println(total)
	return total / float64(len(numbers))
}
