package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()
	// 'done' never receives true since 'c' is yet not receiving
	<-done
	<-done
	close(c)

	for n := range c {
		fmt.Println(n)
	}
}

// go run -race main.go
// vs
// go run main.go
