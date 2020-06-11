package main

/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names. Each line of the text file has
a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name,
and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively
read each line of the text file and create a struct which contains the first and last names found
in the file. Each struct created will be added to a slice, and after all lines have been read from
the file, your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through
your slice of structs and print the first and last names found in each struct.
*/

import (
	"fmt"
	"io/ioutil"
	"log"
	. "strings"
)

func main() {

	type Person  struct {
		fname string
		lname string 
	}	

	group := []Person{} 

	dat, err := ioutil.ReadFile("D:\\Code\\go\\names.txt")
	if err != nil {
		fmt.Println(err)
		log.Panic(err)		
	}
	
    for _ , v := range Split(TrimSpace(string(dat)), "\n") {
		qq := Split(v," ")		
		p:= Person{qq[0], qq[1]}
		group = append(group, p)	
	}

	//fmt.Println(group)
	for i , p:= range group {
		fmt.Println(i, p.fname, p.lname)
	}

}
