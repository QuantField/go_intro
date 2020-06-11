package main

import "fmt"


func sum (nums... int) int {
	s:=0
for _, n := range nums {
s += n
}
	return s
}

func main() {
    fmt.Println(sum(1,2,3,4,5))

}