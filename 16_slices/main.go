package main

import (
	"fmt"
)

func main() {
	myIntSlice := []int{1, 3, 4, 5, 5, 7}
	fmt.Println(myIntSlice)

	myStingSlice := []string{"a", "b", "c", "d", "e"}
	fmt.Println(myStingSlice)
	fmt.Println(myStingSlice[2:4])
	fmt.Println(myStingSlice[2])
	fmt.Println("mySlice", myStingSlice)
	fmt.Println("mySlice"[2])
	fmt.Println([]byte("mySlice"))
	// appendSlice(myStingSlice)
	// makeSlice(5)
	// deleteFromSlice(2, myStingSlice)
	// sliceExample1()

	// sliceInSlice()
	incrementSlice()
}

func deleteFromSlice(i int, slice []string) {
	fmt.Println(append(slice[:i], slice[i+1:]...))
}

func appendSlice(slice []string) {
	newSlice := []string{"d", "g"}
	newSlice = append(newSlice, slice...)
	fmt.Println(newSlice)
}
func makeSlice(x int) {
	slice := make([]int, x)
	fmt.Println(slice)

	student := make([]string, 35)
	students := make([][]string, 35)
	fmt.Println("make")
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)

	var student1 []string
	var students1 [][]string
	fmt.Println("var")
	fmt.Println(student1)
	fmt.Println(students1)
	fmt.Println(student1 == nil)

	student2 := []string{}
	students2 := [][]string{}
	fmt.Println("short hand")
	fmt.Println(student2)
	fmt.Println(students2)
	fmt.Println(student2 == nil)

}

func sliceExample1() {
	mySlice := make([]int, 0, 5)
	fmt.Println("----------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("----------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "capacity: ", cap(mySlice), "Value: ", mySlice[i])
	}
}

func sliceInSlice() {
	records := make([][]string, 0)
	// student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathanr"
	student1[2] = "100.00"
	student1[3] = "74.00"
	// store the record
	records = append(records, student1)
	// student 2
	student2 := []string{"Gomez", "Lisa", "92.00", "96.00"}
	// student2[0] = "Gomez"
	// student2[1] = "Lisa"
	// student2[2] = "92.00"
	// student2[3] = "96.00"
	// store the record
	records = append(records, student2)
	fmt.Println(records)

	transactions := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		transaction := []int{}
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)

}

func incrementSlice() {
	mySlice := make([]int, 1)
	fmt.Println(mySlice[0])
	mySlice[0] = 7
	mySlice = append(mySlice, 7)
	fmt.Println(mySlice[0])
	mySlice[0]++
	fmt.Println(mySlice[0])

}
