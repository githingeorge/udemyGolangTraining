package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(greet("Githin", "George"))
	fmt.Println(multiGreet("Githin", "George"))
}

func greet(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

func multiGreet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
