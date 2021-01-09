package main

import "fmt"

func main() {
	var a []int
	fmt.Println(a)
	fmt.Println("Length: ", len(a))
	fmt.Println("Capacity: ", cap(a))

	a = append(a, 1, 2, 3, 4, 5)
	fmt.Println(a)
	fmt.Println("Length: ", len(a))
	fmt.Println("Capacity: ", cap(a))

	a = append(a, []int{6, 7, 8}...)
	fmt.Println(a)
	fmt.Println("Length: ", len(a))
	fmt.Println("Capacity: ", cap(a))

	x := make([]int, 10, 100)
	x[0] = 42
	x[9] = 999
	//x[10] = 10999 //index out of range [10] with length 10
	x = append(x, 10999) //works fine
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
