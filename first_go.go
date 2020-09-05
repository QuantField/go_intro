package main

import "fmt"

// equivalent of enum
const (
	monday    = iota
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

func main() {

	numbers := []int{2, 3, 4, 5}
	doubles := []float64{2.3, 45.6, 100.3}

	for i, v := range numbers {
		fmt.Println(i, v)
	}

	fmt.Println()
	for i := 0; i < len(doubles); i++ {
		fmt.Println(i, doubles[i])
	}

	//fmt.Println("hello world")
    fmt.Println()
	fmt.Println("Monday:",monday)
	fmt.Println("Sunday:",sunday)
}
