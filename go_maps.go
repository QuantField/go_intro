

package main 

import (
	"fmt"
)


func main() {

	// method 1
	map1 := make(map[string]int)
	map1["xx"] = 1
	map1["yy"] = 2

	// method 2
	var map2 map[string]int
	map2 = make(map[string]int)
	map2["red"] = 3
	map2["blue"] = 4
	map2["red"] = 14

	// method 3
	map3 := map[string]int{
		"one":   1,
		"two":   2,
        "three": 3,
	} 
	
	// method 3
	
	type point struct{
		x, y float32
	}

	map4 := map[string]point{
		"one":   {0, 0},
		"two":   {1, 1},
		"three": {2, 2},
	} 

		
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
	fmt.Println(map4)

	fmt.Println("----------------------------")
	//An empty interface may hold values of any type. 
	//(Every type implements at least zero methods.)
    //Empty interfaces are used by code that handles values of unknown type.
	var x interface{}
	x = 12
	fmt.Println(x)

	x = "hello"
	fmt.Println(x)

	mapx := map[string]interface{}{
		"key1": 123,
		"key2": "A string",
		"key3": []interface{}{3, "Yeah", 23.2},
	}
	fmt.Println(mapx)
	fmt.Println(mapx["key3"])

}
