package main

import "fmt"

func main() {
	var i float32
	var j float64

	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)

	n := 3.14
	n = 13.7e72
	n = 2.1e14
	fmt.Printf("%v, %T\n", n, n)
}
