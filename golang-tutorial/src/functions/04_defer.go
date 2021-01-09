package main

import "fmt"

func main() {
	defer hop()
	bar()
}

func hop() {
	fmt.Println("hop")
}

func bar() {
	fmt.Println("bar")
}
