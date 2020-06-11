package main

import (
	"fmt"
)

func main() {

	map1 := map[string]interface{}{
		"A": 1,
		"t": "Hi",
		"q": "How",
	}

	fmt.Println(map1)

	map2 := make(map[string]interface{})
	map2["aa"] = 1
	map2["bb"] = []interface{}{2, 3, "Kak"}
	map2["cc"] = struct {
		x int
		y int
	}{3, 4}

	fmt.Println(map2)

	val := map2["bb"]
	fmt.Println(val)

}
