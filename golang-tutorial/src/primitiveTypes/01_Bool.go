package main

import "fmt"

func main() {
	var b bool
	fmt.Println(b)

	b = true
	fmt.Printf("%v, %T\n", b, b)

	result := 1 == 1
	fmt.Println(result)
}
