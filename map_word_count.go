package main

import (
	"fmt"
	. "strings"
)

// func WordCount(s string) map[string]int {
// 	return map[string]int{"x": 1}
// }

func main() {

	count := make(map[string]int)
	s := "This is one test and is not another this"
	m := Split(ToLower(s), " ")
	for _, word := range m {
		if _, ok := count[word]; ok {
			count[word] += 1
		} else {
			count[word] = 1
		}

	}
	fmt.Println(count)

}
