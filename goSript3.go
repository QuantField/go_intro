package main

import (
	"fmt"
	"math"
        "strconv"
)

type Complex struct{
    real float64 
    imag float64 
}
// equivalent of a constructor
func newComplex(x, y  float64) *Complex {
    zpt := new(Complex)
    zpt.real = x
    zpt.imag = y
    return zpt
}

func (z Complex) modulus() float64 {
    return math.Sqrt(z.real*z.real + z.imag*z.imag)
}

// by creating String() method we are impementing Stringer interface
// make this accessible by fmt.Prinf Println
func (z Complex) String() string {
    return strconv.FormatFloat(z.real, 'g', 6, 64) + " + " +
           strconv.FormatFloat(z.imag, 'g', 6, 64) + "i"       
}

func main(){
        z := new(Complex)
        z.real = 1.0
        z.imag = 2.0
        fmt.Println("modulus = ", z.modulus())
   	fmt.Println("done")
        
        z2 := newComplex(1.0,2.0)
        fmt.Println("modulus = ", z2.modulus())
        fmt.Printf("%s \n", z2)        
   	fmt.Println("done")             
}
