package main

import (
	"fmt"
)

func main() {
	var name interface{} = "Sydney"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Println("value is not a string")
	}

	var val interface{} = 7
	fmt.Printf("%T\n", val)
	fmt.Println(val.(int) + 6)
}
