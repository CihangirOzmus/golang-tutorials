package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

//any other type which has speak() is a human type
//a value can be of more than one type
type human interface {
	speak()
}

func main() {
	p1 := person{"Cihangir", "Ozmus"}
	p1.speak()
	bar1(p1)
}

func (p person) speak() {
	fmt.Printf("%s speaks english.\n", p.firstName)
}

func bar1(h human) {
	fmt.Println("I called human")
}
