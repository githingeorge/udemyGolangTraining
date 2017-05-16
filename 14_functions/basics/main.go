package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	greet("Githin", "George")
	fmt.Printf("%T  - %+v", greet, greet)
}

func greet(fname, lname string) {
	fmt.Println(fname + lname)
}
