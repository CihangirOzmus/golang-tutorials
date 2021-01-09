package main

import "fmt"

func main() {
	a := incrementer()
	b := incrementer()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementer() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
