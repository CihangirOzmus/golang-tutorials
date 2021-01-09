package main

import "fmt"

func main() {
	// 1st way
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
	fmt.Printf("%T\n", c)
}
