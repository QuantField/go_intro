package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	u := mat.NewVecDense(3, []float64{1, 2, 3})

	fmt.Println(*u)

}
