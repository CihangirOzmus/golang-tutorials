package main

import "fmt"

func main() {
	foo2()

	func() {
		fmt.Println("Anonymous function is called.")
	}()

	func(x int) {
		fmt.Println(x)
	}(42)
}

func foo2() {
	fmt.Println("foo2")
}
