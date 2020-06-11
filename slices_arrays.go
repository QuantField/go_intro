

package main 

import (
	"fmt"
)



func main() {
  
  // array variable is the whole array, not a pointer to the first element                        	
  array1:= [3]int{1,2,3}        // array, 
  array2:= [...]int{4,5,6,7,8} // array,

  // slice variable is a pointer to the first element, always pointers  
  slice1:= []int{45,6,7,8}      
  slice2 := make([]int, 10, 15) // length 10, capacity 15 
  slice2[0] = 23
  slice2 = append(slice2, 12)
  slice_from_array := array2[3:5]
  

  fmt.Println(array1)
  fmt.Println(array2)
  fmt.Println(slice1)
  fmt.Println(slice2)
  fmt.Println(slice_from_array, len(slice_from_array), cap(slice_from_array))


   
}
