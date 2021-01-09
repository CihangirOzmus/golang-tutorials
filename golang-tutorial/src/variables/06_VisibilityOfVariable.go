package main

import "fmt"

var I int = 25 //PascalCase global level
var j int = 27 //camelCase package level

func main() {
	//there is no private scope, use block scope
	var k int = 29 //block scope
	fmt.Println(I, j, k)
}
