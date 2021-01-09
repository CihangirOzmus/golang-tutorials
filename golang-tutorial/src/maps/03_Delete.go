package main

import "fmt"

func main() {
	m := map[string]int{
		"james": 32,
		"jones": 27,
	}

	delete(m, "jones")
	fmt.Println(m)

	if v, ok := m["Jimmy"]; ok {
		fmt.Println(v)
		delete(m, "Jimmy")
	}
}
