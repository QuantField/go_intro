package main

import (
	"fmt"
)

func main() {

	var indexRuneTests = []struct {
		s    string
		rune rune
		out  int
	}{
		{"a A x", 'A', 2},
		{"some_text=some_value", '=', 9},
		{"☺a", 'a', 3},
		{"a☻☺b", '☺', 4},
	}

	var map1 = map[int]struct {
		s string
		p string
	}{
		1:  {"A", "DD"},
		2:  {"BB", "lte"},
		10: {"CC", "qde"},
	}

	fmt.Println(indexRuneTests)
	fmt.Println(indexRuneTests[0])
	fmt.Println(map1)
	fmt.Println(map1[2].s)
	fmt.Println(map1[2].p)

}
