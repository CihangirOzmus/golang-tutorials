package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo // no need explicit declaration
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Boron",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 123456,
		},
	}

	jim.print()
	jim.updateName("Cihangir")
	jim.print()
}

func (p *person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(name string) {
	p.firstName = name
}
