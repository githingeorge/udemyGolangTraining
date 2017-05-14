package main

import (
	"fmt"

	"github.com/githingeorge/udemyGolangTraining/04_scope/vis"
)

var x int = 42

func main() {
	fmt.Println(x)
	foo()
	y := 45
	fmt.Println(y)
	fmt.Println(z)

	max := max(32)
	fmt.Println(max)

	fmt.Println(vis.MyName)
	vis.PrintVar()

	fmt.Println(increment())
	fmt.Println(increment())

	newIncrementer := wrapper()

	fmt.Println(newIncrementer())
	fmt.Println(newIncrementer())
}

var z int = 48

func foo() {
	fmt.Println(x)
	// fmt.Println(y)
}
func max(x int) int {
	return 42 + x

}

var a = 0

func increment() int {
	a++
	return a
}

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x

	}
}
