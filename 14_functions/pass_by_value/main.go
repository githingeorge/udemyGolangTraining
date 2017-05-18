package main

import (
	"fmt"
)

func githin() {
	fmt.Println("githin")
}

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}
func main() {
	age := 42
	fmt.Println(&age)

	changeMe(&age)
	fmt.Println(&age)
	fmt.Println(age)

}
func changeMe(z *int) {
	fmt.Println(z)
	fmt.Println(*z)
	*z = 24
	fmt.Println(z)
	fmt.Println(*z)
}
