package main

import (
	"fmt"
	"math"
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
	c := circle{5}
	info(c)
	c = c.setRadius(10).setRadius(8)
	fmt.Println("-----")
	info(c)
}
