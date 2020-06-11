package main

import "fmt"

// compile and run go run script.go
func square(x float64) float64 {
	return x * x
}

func minmax(x int, y int) (int, int) {
	min := x
	max := y
	if min > max {
		max = x
		min = y
	}

	return min, max
}

func main() {
	S := 0.0
	var i float64 = 1.0
	for i = 1; i <= 1000000.0; i++ {
		S += i
	}
	fmt.Println("Hello, world", S)
	fmt.Println("squre of 3 = ", square(3))
	a, b := minmax(10, 8)
	fmt.Println(a, b)
}
