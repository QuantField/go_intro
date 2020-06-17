package main

import (
	"fmt"
	"math"
	"math/rand"
)

// using channels, don't have to worry about waiting

func calcPi(N int, result chan int) {
	P := 0
	for k := 1; k <= N; k++ {
		x := rand.Float64()
		y := rand.Float64()
		if math.Sqrt(x*x+y*y) <= 1 {
			P++
		}
	}
	// sending to channel result
	result <- P
}

func main() {
	const N = 10_000_000
	nP := 5
	inCirle := make(chan int)
	for i := 0; i < nP; i++ {
		go calcPi(N, inCirle)
	}

	sum := 0
	for i := 0; i < nP; i++ {
		// receiving form channel
		sum += <-inCirle
	}
	pi := 4 * float64(sum) / float64(nP*N)

	fmt.Println(pi)

}
