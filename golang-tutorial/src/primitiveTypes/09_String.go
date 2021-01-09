package main

import "fmt"

func main() {
	s := "This is a string" //string is UTF8
	fmt.Printf("%v, %T\n", s, s)
	fmt.Println(s[6]) //strings are alias for bytes

	s2 := " this is also string"
	fmt.Println(s + s2)

	b := []byte(s)
	fmt.Printf("%v, %T\n", b, b)
}
