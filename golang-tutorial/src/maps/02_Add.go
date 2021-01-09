package main

import "fmt"

func main() {
	m := map[string]int{
		"james": 32,
		"jones": 27,
	}

	m["jane"] = 25

	fmt.Println(m)
}
