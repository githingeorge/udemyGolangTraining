package main

import (
	"fmt"
)

func main() {
	x := "githin"
	switch {
	case len(x) > 1:
		fmt.Println("Githin")
		// fallthrough

	case len(x) < 1:
		fmt.Println(" George")
	default:
		fmt.Println("None")

	}

	switchOnType(7)
	switchOnType("Githin")
	switchOnType(Contact{"githin", "Hello"})
}

type Contact struct {
	name     string
	greeting string
}

func switchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")

	case Contact:
		fmt.Println("Contact")
	default:
		fmt.Println("unKnown")

	}

}
