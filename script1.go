package main

import "fmt"

type person struct {
	name string
	surname string
}

func newPerson(name, surname string) *person {
	p := new(person)
	p.name, p.surname = name, surname
	return p 
}

func (p person) speak(){
	fmt.Println("I am ", p.name, p.surname)
}

// kind of extending class
type worker struct {
	person
	job string
}
//over write speak()
func (w worker) speak(){
	fmt.Println("I am ", w.name, w.surname, "my job is :", w.job)
}
//
/*
func newWorker(name, surname, job string) *worker {
	w := new(worker)
	w.person := newPerson(name, surname)
	w.job = job 
	return w 
}
*/

type canSpeak interface {
	speak()
}

func talk( a canSpeak) {
	a.speak()
}

func main(){
	p:= newPerson("Al", "Jean")
	p.speak()	
	p2 := person{"Ding","Dong"}
	p2.speak()
	p3 := worker{person{"Ding","Dong"},"Developer"  }
	// if speak is not overwritten person.speak() will be used.
	p3.speak()	
	fmt.Println(p3.person.name)
	fmt.Println(p3.name)
	
	talk(p)
	talk(p3)		
}