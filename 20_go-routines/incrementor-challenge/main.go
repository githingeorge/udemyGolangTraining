package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := incrementor("1")
	ch2 := incrementor("2")
	count := 0
	for str := range merge(ch1, ch2) {
		count++
		fmt.Println(str)

	}
	fmt.Println("Counter: ", count)
}

func incrementor(s string) chan string {
	out := make(chan string)
	go func(s string) {

		for i := 0; i < 20; i++ {
			out <- fmt.Sprintln("Process: "+s+" printing: ", i)
			time.Sleep(time.Millisecond + time.Duration(rand.Intn(100)))
		}
		close(out)
	}(s)

	return out
}

func merge(chs ...chan string) chan string {
	out := make(chan string)
	done := make(chan bool)
	for _, ch := range chs {
		go func(in chan string) {
			for s := range in {
				out <- s
			}
			done <- true

		}(ch)
	}

	go func() {
		for i := 0; i < len(chs); i++ {
			<-done

		}
		close(out)
	}()

	return out
}
