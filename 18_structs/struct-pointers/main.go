package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {

	p1 := &person{"James Bond", 20}
	fmt.Println(p1)
	fmt.Printf("%T %v \n", p1, p1)

	fmt.Println(p1.name)
	fmt.Println(p1.age)

}
