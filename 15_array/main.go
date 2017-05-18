package main

import (
	"fmt"
)

func main() {
	var x [58]string
	if len(x) == 0 {
		fmt.Println("x is nil")
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)

	fmt.Println(len(x))
	fmt.Println(x[42])

	var y [256]byte
	fmt.Println(len(y))
	fmt.Println(y[42])
	for i, _ := range y {
		y[i] = byte(i)
	}
	for i, v := range y {
		fmt.Printf("%T - %v - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}

}
