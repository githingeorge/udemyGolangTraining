package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		var msg string

		if i%3 == 0 {
			msg = "Fizz"
		}
		if i%5 == 0 {
			msg += "Buzz"

		}
		if msg == "" {
			fmt.Println(i)
		} else {

			fmt.Println(i, " -- ", msg)
		}

	}
}
