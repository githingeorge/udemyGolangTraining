package main

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

func (p Person) Greeting() {
	fmt.Println("I'm just a reqular pereson.")
}

type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

func (dz DoubleZero) Greeting() {
	fmt.Println("Miss MoneyPenny, so good to see you")
}

func main() {
	p := Person{
		First: "Ian Flemning",
		Age:   44,
	}
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}
	// p2 := DoubleZero{
	// 	Person: Person{
	// 		First: "Miss",
	// 		Last:  "MoneyPenny",
	// 		Age:   19,
	// 	},
	// 	First:         "If looks could kill",
	// 	LicenseToKill: false,
	// }
	// // fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill, p1.Person.First)
	// fmt.Printf("%T - %v \n", p1, p1)

	// fmt.Println(p2.First, p2.Last, p2.Age, p2.LicenseToKill, p2.Person.First)
	// fmt.Printf("%T - %v \n", p2, p2)

	p.Greeting()
	p1.Greeting()
	p1.Person.Greeting()
}
