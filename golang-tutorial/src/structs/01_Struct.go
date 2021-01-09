package main

import "fmt"

type person struct {
	first   string
	last    string
	age     int
	address address
}

type address struct {
	street  string
	country string
}

func main() {
	p1 := person{"james", "bond", 32, address{"1st street", "london"}}
	p2 := person{"jane", "penny", 27, address{"2d street", "london"}}
	p3 := person{first: "michael", last: "bondy", age: 37, address: address{street: "3rd street", country: "london"}}

	fmt.Println(p1, p2, p3)
	fmt.Println(p1.first)
	fmt.Println(p1.address.country)
}
