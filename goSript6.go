package main 

import (
    "fmt"
    "math/rand"
    "math"
	"time"
)



func calcPi(N int) float64{
    P := 0
    for i:=1; i<=N; i++ {
        x:= rand.Float64()
        y:= rand.Float64()
        if (math.Sqrt(x*x+y*y)<=1) { P++ }        
    }
    pi := 4*float64(P)/float64(N)
	fmt.Println(pi)
	return pi
}


func main() {
  N:=1000000
  go calcPi(N)
  go calcPi(N)
  go calcPi(N)
  go calcPi(N)
  go calcPi(N)
  go calcPi(N)
  go calcPi(N)
  go calcPi(N)
  time.Sleep(5000* time.Millisecond)   
}


