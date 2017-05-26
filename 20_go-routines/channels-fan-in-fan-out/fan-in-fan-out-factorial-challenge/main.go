package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers_ch := genNumbers()

	// f1 := factorial(numbers_ch)
	// f2 := factorial(numbers_ch)

	// for n := range merge(f1, f2) {
	// 	fmt.Println("factorial: ", n)
	// }
	chSlice := fanOut(numbers_ch, 10)

	var y int
	for n := range merge(chSlice...) {
		y++
		fmt.Println(y, "factorial: ", n)
	}
}

func genNumbers() chan int {
	out := make(chan int)
	go func() {

		for i := 0; i < 2; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fanOut(in chan int, n int) []chan int {

	chSlice := make([]chan int, n)
	fmt.Println(chSlice)
	for i := 0; i < n; i++ {
		chSlice[i] = factorial(in)
	}
	fmt.Println(chSlice)

	return chSlice
}

func factorial(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			// fmt.Println(n)
			out <- fact(n)
		}
		close(out)
	}()
	return out

}

func fact(n int) int {
	factorial := 1
	i := int(n)
	for ; i > 1; i-- {
		factorial *= i
	}
	return factorial
}

func merge(ins ...chan int) chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	wg.Add(len(ins))

	for _, ch := range ins {
		go func(in chan int) {
			for n := range in {
				out <- n
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
