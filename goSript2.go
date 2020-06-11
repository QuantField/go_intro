package main

import "fmt"
import "math"


type realFunction func(float64) float64


type double float64


const eps=1.0E-6

func  derivative( f realFunction, x float64) float64 {
    return (f(x+eps) - f(x-eps))/(2*eps)
}

func sq( x float64) float64 {
    return x*x - 2
}


func newtonRaphson( g realFunction, x float64  ) (float64, int) {
    xnew  := x
    xold  := xnew+0.1
    var iter int = 0
    for math.Abs((xnew-xold)/xnew)>=eps {
        iter++
        xold = xnew
	xnew  = xnew - g(xnew)/derivative(g,xnew)
        fmt.Println(xnew)
    }   
    return xnew, iter
}


//---- Example to attaching function to types (similar to oop)
func ( c double) Sqrt() float64 {
    return math.Sqrt(float64(c)) // must cast back to float64
}



func main() {
    
    sol, iter := newtonRaphson(sq, 0.5) 
    fmt.Println("Solultion = ", sol, "  iterations = ", iter)
    fmt.Println()
    
    s:= double(2.0)
    fmt.Println("Square root :", s.Sqrt())
    
    fmt.Println("")
    
    cub := func(x float64) float64 { return x*x*x - 8 }
    
    newtonRaphson(cub, 0.5)    
    
    newtonRaphson(func(x float64) float64 { return x*x*x - 8 }, 1)    
    
}  
    