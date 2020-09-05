package main

import (
	"fmt"
)

func main() {

	list := []interface{}{2, 3, "Hello", 2.32, 4}

	list2 := make([]interface{}, 10, 15)
	list2[0] = 3.14
	list2[3] = "This is a test"

	fmt.Println(list)
	fmt.Println(list[1])
	fmt.Println(list2)

	var listx []interface{}
	listx = append(listx, 34.5)
	listx = append(listx, "Test")
	listx = append(listx, "a", 0, 23)
	fmt.Println(listx)

}
