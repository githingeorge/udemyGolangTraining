package main

import (
	"fmt"
)

func main() {
	greeting := func() {
		fmt.Println("hello world!")
	}
	greeting()
	greeter := makeGreeter()
	fmt.Println(greeter())
	fmt.Println(makeGreeter()())

}

func makeGreeter() func() string {
	return func() string {
		return "Hello world from greeter!"
	}
}
