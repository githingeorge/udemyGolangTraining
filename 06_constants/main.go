package main

import (
	"fmt"
)

const p = "death and taxes"

const (
	a = iota
	b = iota
	c = iota
	d = iota
	e = iota
)

// fmt.Println(a, b, c, d, e)

const (
	cold = "COLD"
	hot  = "HOT"
)

const (
	_  = iota
	qa = iota * 10
	qb = iota * 20
)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 20)
)

func main() {
	const q = 32
	fmt.Println(a, b, c, d, e)
	fmt.Println(qa, qb)
	fmt.Println(KB, MB)

	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(cold)
	fmt.Println(hot)

}
