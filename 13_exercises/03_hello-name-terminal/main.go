package main

import (
	"fmt"
)

func main() {
	fmt.Println("whats your name?")
	var name string

	fmt.Scan(&name)
	fmt.Printf("Hello %s", name)
}
