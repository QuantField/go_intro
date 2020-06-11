package main

import (
    "fmt"
)

type realFunction func(float64) float64 

func tabulateFunction(f realFunction, x []float64) []float64 {
    N := len(x)
    y := make([]float64,N)
    for i,_ := range x {
        y[i] = f(x[i])
    }
    return y
}
// instead of using realFunction we can just use the whole thing.
func tabulateFunction2(f func(float64) float64, x []float64 ) []float64 {
    N := len(x)
    y := make([]float64,N)
    for i,_ := range x {
        y[i] = f(x[i])
    }
    return y
}

func trapezoid(y []float64, h float64) float64 {
    S :=0.0
    for i:=1; i<len(y)-1; i++{
        S+=y[i]
    }
    S = S*h + (y[0]+y[len(y)-1])*h/2.0
    return S
}

func simpson(y []float64, h float64) float64 {
    n := len(y)-1
    S := y[0]+y[n]   
    for i:=1; i<=n-1; i=i+2 { S+= 4.0*y[i] }  
    for i:=2; i<=n-2; i=i+2 { S+= 2.0*y[i] }
    return S*h/3.0
}

// returning fucntions 
func geneQuad( c float64) realFunction {
    f := func(x float64) float64 { return c*x*x }
    return f
}


func main(){
    const N = 10.0 
    xa, xb := 0.0, 1.0
    h :=(xb-xa)/N
    
    a:= make([]float64,N+1)
    
    for i, _ := range a {
        a[i] = xa+float64(i)*h
        //fmt.Println(a[i])
    }
    
    g := func(x float64) float64 { return x*x }
    
    //s := tabulateFunction(g, a)  
    s := tabulateFunction2(g, a)  
    
    //for  _, v := range s {
    //    fmt.Println(v)
    //}
    
    Int := trapezoid(s, h)
    Int2:= simpson(s, h)
    
    fmt.Println(" Trapezoid : ", Int)
    fmt.Println(" Simpson   : ", Int2)
    
    d := geneQuad(0.5)
    fmt.Println(" g(x)   : ", d(2))
    
    T :=1
    fmt.Println( int(T==1))
       
    
}