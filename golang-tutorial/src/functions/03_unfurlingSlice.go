package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("sum = ", sum(x...))
}

func sum(x ...int) int {
	fmt.Println(x)
	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}
