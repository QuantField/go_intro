package main

/*
Methods in go are function attached to struct, similar to methodes in OOP
*/

import (
	"fmt"
	. "math"
	// "io/ioutil"
	// "log"
	// . "strings"
)


type shape interface {
	surface() float64
	volume()  float64
}

//----------------------------- 

type sphere struct {
	radius float64
}

func (s *sphere) volume() float64 {
	return 4*Pi*Pow(s.radius,3)/3
}

func (s *sphere) surface() float64 {
	return 4*Pi*Pow(s.radius,2)
}

//----------------------------- 

type cube struct {
	side float64
}

func (s *cube) volume() float64 {
	return Pow(s.side,3)
}

func (s *cube) surface() float64 {
	return s.side*s.side
}

//----------------------------------------

func print_area_volume( sh shape) {
	fmt.Println("Area   = ", sh.surface())
	fmt.Println("Volume = ", sh.volume())
}

func main() {

	sph  := sphere{10}
	cb   := cube{side:2} // also valid specifying the field name

	fmt.Println(sph.volume(), sph.surface())
	fmt.Println(cb.volume(), cb.surface())
	fmt.Println()
	print_area_volume(&sph)

	
}
