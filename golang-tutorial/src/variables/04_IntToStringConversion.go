package main

import (
	"fmt"
	"strconv"
)

//use strconv package for int to string conversion
func main() {
	var i int = 42

	var j string
	j = string(i)
	var k = strconv.Itoa(i)

	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", k, k)
}
