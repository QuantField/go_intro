package main

import (
	"fmt"
	"log"
)

// not working

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	res := 1 / 0
	fmt.Println(res)

}
