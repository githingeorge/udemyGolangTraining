package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 40; i++ {
		fmt.Println("foo: ", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 40; i++ {
		fmt.Println("bar: ", i)
		time.Sleep(time.Duration(20 * time.Millisecond))
	}
	wg.Done()
}
