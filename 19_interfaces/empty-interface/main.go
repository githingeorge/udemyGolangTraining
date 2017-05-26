package main

import (
	"fmt"
)

type vehicles interface{}

type Vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type Car struct {
	Vehicle
	Wheels int
	Doors  int
}

type Plane struct {
	Vehicle
	Jet bool
}

type Boat struct {
	Vehicle
	length int
}

func main() {
	dzire := Car{}
	ritz := Car{}
	alto := Car{}

	boeing747 := Plane{}
	boeing757 := Plane{}
	boeing767 := Plane{}

	sanger := Boat{}
	noutique := Boat{}
	malibu := Boat{}

	rides := []vehicles{dzire, ritz, alto, boeing747, boeing757, boeing767, sanger, noutique, malibu}
	for _, ride := range rides {
		fmt.Println(ride)

	}
}
