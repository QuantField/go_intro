package main

/*
Use of sync to wait for all routines to terminates

*/

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func f(n int) {
	for i := 0; i < 5; i++ {
		fmt.Println("thread ", n, ":", i)
	}
	wg.Done()
}
func main() {
	// runing 10 go routines
	for i := 0; i < 10; i++ {
		wg.Add(1) // adding goroutine
		go f(i)
	}
	wg.Wait()
}
