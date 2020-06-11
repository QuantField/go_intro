package main

/*
Methods in go are function attached to struct, similar to methodes in OOP
*/

import (
	"fmt"
	// "io/ioutil"
	// "log"
	// . "strings"
)

//----------------------- Type record and its methods -----------------------
type record struct {
	name string
	age  int 
}

// here it is better to use references, i.e. r *record, as the default is 
// passing by values.
func (r *record) display(){
	fmt.Println("Name: ", r.name)
	fmt.Println("Age: ", r.age)
}

// important to use reference *record, if not the change is local and won't 
// have any effect
func (r *record) set(name string, age int) {
	r.name = name 
	r.age = age 
}

//---------------------------------------------------------------------------


func main() {
	p := record{"Kali Horn", 25}
	p.display()
	p.set("New Guy", 50)
	p.display()
	
}
