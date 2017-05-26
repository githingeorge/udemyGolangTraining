package main

import (
	"fmt"
)

func main() {
	// makeMap()
	// deleteFromMap()
	// checkKeyInMap()

	rangeOverMap()

}

func makeMap() {
	myGreeting := make(map[string]string)
	myGreeting["Tim"] = "Good Morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)

	myGreeting1 := map[string]string{
		"Tim":   "Good Morning.",
		"Jenny": "Bonjour.",
	}

	fmt.Println(myGreeting1)

}

func deleteFromMap() {
	myGreeting := map[string]string{
		"kalabha": "Stupid girl",
		"Shilpa":  "Dark singing girl",
		"Remya":   "Poisonus girl",
		"Anushma": "Endosulfan girl",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, "kalabha")
	fmt.Println(myGreeting)

}

func checkKeyInMap() {
	myGreeting := map[string]string{
		"kalabha": "Stupid girl",
		"Shilpa":  "Dark singing girl",
		"Remya":   "Poisonus girl",
		"Anushma": "Endosulfan girl",
	}

	fmt.Println(myGreeting)
	if val, exists := myGreeting["kalabha"]; exists {
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}
	delete(myGreeting, "kalabha")
	if val, exists := myGreeting["kalabha"]; exists {
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}
	fmt.Println(myGreeting)
	delete(myGreeting, "kalabha")

}

func rangeOverMap() {
	myGreeting := map[int]string{
		0: "Good MOrning!",
		1: "Bonjour",
		2: "Buenous dias",
		3: "Bongiorno!",
	}

	for i, v := range myGreeting {
		fmt.Println(i, " - ", v)
	}

}
