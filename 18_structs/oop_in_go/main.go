package main

import (
	"fmt"
)

type foo int

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}
	fmt.Println(p1.First, p1.Last, p1.Age, p1.LicenseToKill, p1.Person.First)
	fmt.Printf("%T - %v \n", p1, p1)
}
