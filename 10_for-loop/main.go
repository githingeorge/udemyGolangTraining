package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, "-", j)
		}
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	i = 0
	for {
		i++
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
		if i > 10 {
			break
		}
	}

	for i := 5000; i < 5100; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

}
