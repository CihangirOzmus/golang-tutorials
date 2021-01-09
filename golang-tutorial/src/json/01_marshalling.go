package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p1 := person{"James", "Bond", 42}
	p, err := json.Marshal(p1)

	if err != nil {
		fmt.Println("Error!")
	}

	fmt.Println(string(p))

	p2 := person{"Jane", "Doe", 32}
	ps := []person{p1, p2}

	pJson, err := json.Marshal(ps)

	if err != nil {
		fmt.Println("Error!")
	}

	fmt.Println(string(pJson))
}
