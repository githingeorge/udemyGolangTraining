package main

import (
	"fmt"
)

type foo int

type person struct {
	first string
	last string
	age int


}

func main() {
	var myAge foo
	myAge = 14
	fmt.Printf("%T %v \n", myAge, myAge)

	p1:=person{"James","Bond", 20}
	p2:=person{"Miss","Moneypenny", 20}
	fmt.Printf("%T %v \n", p1, p1)
	fmt.Printf("%T %v \n", p2, p2)
	
}
