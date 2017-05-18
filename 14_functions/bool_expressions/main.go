package main

import (
	"fmt"
)

func main() {
	if true || false {
		fmt.Println("I'm driving!")
	}

	if true && false {
		fmt.Println("I'm driving!")
	}

}
