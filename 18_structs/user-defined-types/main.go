package main

import (
	"fmt"
)

type foo int

type person struct {
	first string
	last  string
	age   int
}

func main() {
	var myAge foo = 44
	fmt.Printf("%T - %v \n", myAge, myAge)

	var yourAge int = 44
	fmt.Printf("%T - %v \n", yourAge, yourAge)

	fmt.Println(myAge + foo(yourAge))

	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 20}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Printf("%T %v \n", p1, p1)
	fmt.Printf("%T %v \n", p2, p2)

}
