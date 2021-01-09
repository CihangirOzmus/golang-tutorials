package main

import "fmt"

func main() {
	x := returnFunc()
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)
}

func returnFunc() func() int {
	return func() int {
		return 451
	}
}
