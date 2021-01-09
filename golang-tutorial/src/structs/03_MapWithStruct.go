package main

import "fmt"

type employee struct {
	first string
	last  string
	age   int
}

func main() {
	m := map[string]employee{
		"e1": {"james", "bond", 32},
		"e2": {"jane", "penny", 27},
	}

	e3 := employee{"james", "bond", 32}

	m["e3"] = e3
	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v.first, v.last, v.age)
	}
}
