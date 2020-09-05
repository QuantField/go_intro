/*
Q12. (1) Map function
A map()-function is a function that takes a function and a
list. The function is applied to each member in the list and a new list containing these
calculated values is returned. Thus:
map(f(), (a1, a2, . . . , an−1, an)) = (f(a1), f(a2), . . . , f(an−1), f(an))
1. Write a simple map()-function in Go. It is sufficient for this function only to work
for ints.
2. Expand your code to also work on a list of strings
*/

package main 

import (
	"fmt"
)


type int_func func(int) int 

// variadic arguments ...
func map_f( f int_func, a ... int) []int{

	list_f := make([]int, len(a))
	for i , v := range a {
		list_f[i] = f(v)
	}
	return list_f
}


func main() {
    // equivalent python lambda functions
	sqr := func (x int) int { return x*x}

	fmt.Println(  map_f(sqr, 2,3,4,5) )
}
