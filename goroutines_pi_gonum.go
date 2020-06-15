package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

// this is how to get gonum the ... should be there
// go get -u gonum.org/v1/gonum/...

var wg sync.WaitGroup

func calcPi(N int, result *float64) {
	P := 0
	for k := 1; k <= N; k++ {
		x := rand.Float64()
		y := rand.Float64()
		if math.Sqrt(x*x+y*y) <= 1 {
			P++
		}
	}
	*result = 4 * float64(P) / float64(N)
	wg.Done()
}

func main() {
	N := 10_000_000
	const nThreads = 20
	var piArr [nThreads]float64

	for i := 0; i < nThreads; i++ {
		wg.Add(1)
		go calcPi(N, &piArr[i])
	}
	wg.Wait()

	fmt.Println(piArr)

	t := floats.Sum(piArr[:])

	fmt.Println("mean   = ", stat.Mean(piArr[:], nil))
	std := stat.StdDev(piArr[:], nil)
	nSamples := len(piArr)
	fmt.Println("std    = ", std)
	fmt.Println("stderr = ", stat.StdErr(std, float64(nSamples)))
	fmt.Println(t)

	//time.Sleep(2000 * time.Millisecond)
}
