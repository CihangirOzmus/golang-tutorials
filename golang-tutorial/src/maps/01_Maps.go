package main

import "fmt"

func main() {
	m := map[string]int{
		"james": 32,
		"jones": 27,
	}

	fmt.Println(m)
	fmt.Println(m["james"])
	fmt.Println(m["jonathan"])

	v, ok := m["james"]
	fmt.Println(v, ok)

	val, ok := m["jane"]
	fmt.Println(val, ok)

	if v, ok := m["james"]; ok {
		fmt.Println("James is", v, "years old.")
	}

	for i, v := range m {
		fmt.Println(i, v)
	}
}
