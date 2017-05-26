package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)

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

func merge(in1, in2 chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for n := range in1 {
			out <- n
		}
		wg.Done()
	}()

	go func() {
		for n := range in2 {
			out <- n
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
