package main

import (
	"fmt"
	"sort"
)

type People []string

func (p People) Len() int { return len(p) }

func (p People) Less(i, j int) bool { return p[i] < p[j] }

func (p People) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func main() {
	studyGroup := People{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)

	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	fmt.Println("------ with people type")
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)

	sort.Sort(People(s))
	fmt.Println(s)

	fmt.Println("------ with string slice")
	ss := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(ss)

	sort.StringSlice(ss).Sort()
	fmt.Println(ss)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	fmt.Println(n)

	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)

}
