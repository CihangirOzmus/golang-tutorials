package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("func expression is executed.")
	}

	f()

	g := func(x int) {
		fmt.Println(x)
	}

	g(42)
}
