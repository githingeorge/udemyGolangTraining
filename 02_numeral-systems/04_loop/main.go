package main

import (
	"fmt"
)

func main() {
	for index := 60; index < 122; index++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", index, index, index, index)
	}
}
