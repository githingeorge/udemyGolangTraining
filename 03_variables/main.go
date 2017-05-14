package main

import (
	"fmt"
)

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	var e string
	var f int
	var g bool
	var h []int
	var i float64

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)

	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)
	fmt.Printf("%v \n", i)
	fmt.Printf("%T \n", h)

}
