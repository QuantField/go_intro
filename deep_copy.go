package main

import (
	"fmt"
	"sort"
)

func main() {

	list := []int{10, 2, 3, -4}
	//option one
	cp := make([]int, len(list))
	copy(cp, list)
	fmt.Println(list, cp)

	//option two
	cp2 := append([]int{}, list...) // note the ..., unpacking
	fmt.Println(list, cp2)

	// arrays however are copied by deepcopied by default in assignment

	//--- sorting ---
	sort.Ints(list)
	fmt.Println(list)

}
