package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50, 60}
	b := a[:]
	c := a[3:]
	d := a[:3]
	e := a[1:4]

	c[0] = 422
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	//slicing operation uses source, so using array will have the same change
}
