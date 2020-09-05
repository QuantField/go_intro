// function as a struct field

package main

import (
	"fmt"
	"math"
)

// not really an oop approach. Radius is not a method
// as it is set to any function satisfying the defined signature
// when creating the stuct
type Point struct {
	X, Y           float32
	ScaleTranslate func(float32, float32, float32, float32) (float32, float32)
}

type complex struct {
	re, img float64
	module  func(z *complex) float64
}

func main() {

	p := Point{
		X: 1.2,
		Y: 3.4,
		ScaleTranslate: func(scale, displacement, x, y float32) (float32, float32) {
			return scale*x + displacement, scale*y + displacement
		},
	}

	fmt.Println(p)
	fmt.Println(p.ScaleTranslate(1, 2, 0.5, 0.5))

	z := complex{
		re:  0.5,
		img: 1.0,
		module: func(z *complex) float64 {
			// should use &z. but Go does it implicitly
			return math.Sqrt(z.re*z.re + z.img*z.img)
		},
	}

	fmt.Println(z)
	fmt.Println(z.module(&z))

}
