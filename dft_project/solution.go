package main

import (
	"fmt"
	"math"	
)

type Solution struct {
	h float64
	r []float64
	V []float64
	E float64
	K []float64
	Sol []float64	
}


func NewSolution(radius *Radius, V []float64) *Solution{
	N:= len(radius.array)	
	s:= Solution{
		h: radius.h,
		r: radius.array,
		V: V,		
		K: make([]float64, N),
		Sol: make([]float64, N),
	}
	return &s
}


func (s *Solution) Numerov(E float64) {
	    
	N:= len(s.r)
	C:= s.h*s.h/12
	
	for i:=1; i<N ; i++	{
		s.K[i] =  2*(E-s.V[i]) // in Hartree  or E0-2*V[i] for Ryleigs;          
	} 
	
	alph:= math.Sqrt(-2*E)
	s.Sol[N-1] = s.r[N]*math.Exp(-alph*s.r[N-1])
	s.Sol[N-2] = s.r[N-1]*math.Exp(-alph*s.r[N-2])
		
	for i:=N-2; i>=1; i-- { 	    	    		     
		s.Sol[i-1] = ( 2*(1. - 5.*C*s.K[i])*s.Sol[i] - 
				       (1. + C*s.K[i+1] )*s.Sol[i+1] )/(1.+C*s.K[i-1])
	}
}


func main(){

	// defining the ratius with max length of 10 units
	// grid of N+1 points 
	Rmax:= 10.0
	N := 100
	radius := NewRadius(Rmax, N)

	// Atom Z = 1
	atom:= NewAtom(1, radius)

	sol := NewSolution(radius, atom.Veff )

	sol.Numerov(0.5)

	fmt.Println("Solution: ")

	//fmt.Println(sol.Sol)








   
   
  
}