package main

import "errors"
import "log"
import "fmt"

// ErrNorgateMath norgate math: square root of negative number
var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Println(err)
	}
}

// func Sqrt(f float64) (float64, error) {
// 	if f < 0 {
// 		return 0, ErrNorgateMath
// 	}
// 	return 42, nil
// }

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("norgate math: square root of negative number: %v", f)
	}

	return 42, nil
}
