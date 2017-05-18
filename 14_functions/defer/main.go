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
	defer githin()
	defer world()
	hello()
}
