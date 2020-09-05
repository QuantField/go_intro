package main

import (
	"fmt"
)

type Printer interface {
	Print()
}

//----------------------------------------------------
type Person struct {
	name string
	age  int
}

func (p *Person) set(name string, age int) {
	p.name, p.age = name, age
}

func (p *Person) Print() {
	fmt.Println("Name: ", p.name)
	fmt.Println("Age : ", p.age)
}

//----------------------------------------------------
type Employee struct {
	Person
	id     string
	salary float32
}

func CreateEmployee(name string, age int, id string, salary float32) *Employee {
	e := Employee{
		Person: Person{name: name, age: age},
		id:     id,
		salary: salary,
	}
	return &e
}

// func (e *Employee) set(name string, age int , id string, salary float32){
// 	e.name = name
// 	e.age = age
// 	e.id = id
// 	e.salary = salary
// }

func (p *Employee) Print() {
	fmt.Println("Id  : ", p.id)
	fmt.Println("Name: ", p.name)
	fmt.Println("Age : ", p.age)
	fmt.Println("Salary: ", p.salary)
}

func main() {

	p1 := Person{name: "Live OrDie", age: 50}
	fmt.Println(p1)
	p1.set("OLiver", 25)
	fmt.Println(p1)
	p2 := p1
	p2.name = "Kaki"
	fmt.Println(p1.name, p2.name)
	//---------------------------------------
	fmt.Println("--------- Empty object -----------")
	e1 := Employee{}
	fmt.Println(fmt.Printf("e1 = %+v\n", e1))

	fmt.Println("--------- Non Empty object -----------")
	e2 := Employee{
		Person: Person{name: "John Wick", age: 35},
		id:     "124",
		salary: 35000,
	}
	fmt.Println(fmt.Printf("e1 = %+v\n", e2))

	fmt.Println("--------- Using a Constructor -----------")
	e3 := CreateEmployee("Kohl", 45, "223", 55000)
	fmt.Println(e3)

	// e1 := Employee{name:"John Wick",age:50, id:"123", salary:35000}
	// fmt.Println(e1)

	// polymorphysm  invoque Print() accoding to type
	var q Printer // note here both q,w declared as Printer
	var w Printer // because they implicitly implement it via Print()
	q = &p1
	w = &e2
	fmt.Println("---------- interfaces --------------")
	q.Print()
	fmt.Println("----------------")
	w.Print()

	fmt.Println("----- calling Employee's Print() 'super()'------")
	e2.Person.Print()
	fmt.Println("----- calling Employee's Print() ------")
	e2.Print()

}
