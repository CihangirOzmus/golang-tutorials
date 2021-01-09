package main

import "fmt"

func main() {
	a := []int{10, 20, 30}
	fmt.Println(a)
	fmt.Println("length: ", len(a))
	fmt.Println("capacity: ", cap(a))

	b := a //slices are referenced type, no need to use pointer
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}
