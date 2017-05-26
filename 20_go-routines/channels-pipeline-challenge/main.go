package main

import (
	"fmt"
)

func main() {

	c := gen(2, 3, 5, 8, 21, 43)
	out := sq(c)
	for s := range out {
		fmt.Println(s)

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
