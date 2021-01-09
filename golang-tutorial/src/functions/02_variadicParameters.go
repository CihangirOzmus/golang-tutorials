package main

import "fmt"

func main() {
	foo1(1, 2, 3, 4, 5, 6)
}

func foo1(x ...int) {
	fmt.Println(x)

	sum := 0
	for _, v := range x {
		sum += v
	}

	fmt.Println("The total =", sum)
}
