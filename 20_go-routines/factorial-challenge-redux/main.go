package main

import (
	"fmt"
)

func main() {

	for n := range factorial(genNumbers(30)) {
		fmt.Println("factorial: ", n)
	}
}

func genNumbers(n int) chan int {
	out := make(chan int)
	go func() {
		for i := n; i > 0; i-- {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(in chan int) chan uint {
	out := make(chan uint)
	go func() {
		for n := range in {
			// fmt.Println(n)
			total := 1
			for i := n; i > 1; i-- {
				total *= i
			}
			out <- uint(total)
		}
		close(out)
	}()
	return out

}

// func factorial(in chan int) chan int {

// 	go func() {
// 		total := 1
// 		for i := n; i > 0; i-- {
// 			total *= i
// 		}
// 		f <- total
// 		close(f)
// 	}()

// 	return f
// }
