package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1)

	fmt.Println("----------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	fmt.Printf("%T", p1)

}
