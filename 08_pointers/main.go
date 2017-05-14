package main

import (
	"fmt"
)

func main() {
	a := 43
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)

	var b *int = &a
	fmt.Println("b  = ", b, *b)
	*b = 42
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)

	x := 5
	zero(&x)
	fmt.Println("x: ", x)

}

func zero(x *int) {
	*x = 0
}
