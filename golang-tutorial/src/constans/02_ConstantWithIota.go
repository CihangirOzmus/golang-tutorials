package main

import "fmt"

const (
	a = iota
	b = iota
	c = iota
	d
	e
)

//iota belongs to it's block
const (
	a2 = iota
	b2 = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println(a2)
	fmt.Println(b2)
}
