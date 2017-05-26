package main

import (
	"fmt"
	"math"
	"strconv"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	fmt.Println(c)
	return math.Pi * c.radius * c.radius
}

func (c circle) setRadius(r float64) circle {
	c.radius = r
	fmt.Println(c.radius)
	return c
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	var x = 12
	var y = 12.122233
	fmt.Println(y + float64(x))

	var a rune = 'a'
	var b int32 = 'b'
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(string(a))
	fmt.Println(string(b))

	fmt.Println(string([]byte{'h', 'e', 'l', 'l', 'o'}))
	fmt.Println([]byte("hello"))

	var p = "12"
	var q = 6
	z, _ := strconv.Atoi(p)
	fmt.Println(q + z)

	var r = 12
	var s = "I have this many: " + strconv.Itoa(r)
	fmt.Println(s)
	fmt.Printf("I have this many: %v", r)

}
