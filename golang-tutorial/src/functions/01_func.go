package main

import "fmt"

func main() {
	foo()
	fooName("James")
	s := woo("Jane")
	fmt.Println(s)

	s1, s2 := mouse("Jane", "Fleming")
	fmt.Println(s1, s2)
}

func foo() {
	fmt.Println("Hi")
}

func fooName(name string) {
	fmt.Println("Hi", name)
}

func woo(name string) string {
	return fmt.Sprint("Hello from woo ", name)
}

func mouse(s string, l string) (string, bool) {
	a := fmt.Sprint(s, l, ` says "Hello""`)
	b := false
	return a, b
}
