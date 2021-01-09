package main

import (
	"fmt"
)

func main() {
	// 2nd way
	c := make(chan int, 2) //room for 2

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
}
