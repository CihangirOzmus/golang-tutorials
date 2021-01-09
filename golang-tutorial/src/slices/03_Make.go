package main

import "fmt"

func main() {
	a := make([]int, 3)
	fmt.Println(a)
	fmt.Println("Length: ", len(a))
	fmt.Println("Capacity: ", cap(a))

	b := make([]int, 3, 100)
	fmt.Println(b)
	fmt.Println("Length: ", len(b))
	fmt.Println("Capacity: ", cap(b))
}
