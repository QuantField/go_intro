package main

import (
	"fmt"
	"sort"
)

func main() {

	list := []int{10, 2, 3, -4}

	// option 1
	fmt.Println(list)
	sort.Ints(list)
	fmt.Println(list)

	// option 2
	fmt.Println()
	list = []int{10, 2, 3, -4}
	fmt.Println(list)
	//sort.SliceStable keeps original order of equal elements
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	fmt.Println(list)

	fmt.Println()

	people := []struct {
		name string
		age  int
	}{
		{"Helen", 56},
		{"Karim", 21},
		{"Lola", 21},
		{"John", 65},
		{"Karl", 18},
	}
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})
	fmt.Println(people)

	// option 3 implment Less, Swap, Len via interface

}
