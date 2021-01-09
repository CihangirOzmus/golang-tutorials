package main

import "fmt"

//& address
//* value
func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *int = &a
	fmt.Println(b)

	c := &a
	fmt.Println(c)
	fmt.Println(*c)

	*c = 43
	fmt.Println(a)
}
