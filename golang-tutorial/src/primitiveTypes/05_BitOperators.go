package main

import "fmt"

func main() {
	a := 10             //1010
	b := 3              //0011
	fmt.Println(a & b)  //0010
	fmt.Println(a | b)  //1011
	fmt.Println(a ^ b)  //1001
	fmt.Println(a &^ b) //0100
}
