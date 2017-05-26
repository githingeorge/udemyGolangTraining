package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3, 4, 5)

	//  FAN OUT
	// 	Distribute the sq work across 2 goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// FAN IN
	// Consume the merged output from multiple channels.
	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(ins ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {

		go func(in chan int) {
			for n := range in {
				out <- n
			}
			wg.Done()
		}(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
