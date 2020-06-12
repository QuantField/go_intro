package main

import (
	"fmt"
	"strconv"
)

//----------------------------------------------------
type Person struct {
	name string
	age  int
}

// Constructor, by Convention prefix New to struct name
func NewPerson(name string, age int) *Person {
	// add some logic if needed
	p := Person{name: name, age: age}
	return &p
}

// also by convention, receive is first letter of struct name in lower case
func (p *Person) Set(name string, age int) {
	p.name, p.age = name, age
}

// equivalent of Python __str__
func (p *Person) String() string {
	disp := fmt.Sprintf("Person( Name: %s , Age: %s )", p.name, strconv.Itoa(p.age))
	return disp
}

func main() {

	person := NewPerson("Hulk", 45)
	fmt.Println(person) // String() method will be invoqued.

}
