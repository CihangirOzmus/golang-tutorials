package main

import "fmt"

func main() {
	var i float32 = 42.5
	fmt.Printf("%v, %T\n", i, i)

	var j int
	j = int(i)
	fmt.Printf("%v, %T\n", j, j)

	var a int = 10
	var b int8 = 5
	fmt.Println(a + int(b))
}
