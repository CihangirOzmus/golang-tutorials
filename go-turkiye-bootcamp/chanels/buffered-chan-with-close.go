package main

import (
	"fmt"
)

func main() {
	bufferedChan := make(chan int, 5)
	done := make(chan bool)

	go func(chanelName chan int, done chan bool) {
		for val := range chanelName {
			fmt.Println(val)
		}
		fmt.Println("chanel has closed")
		done <- true
	}(bufferedChan, done)

	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3
	bufferedChan <- 4
	bufferedChan <- 5
	bufferedChan <- 6 // after 5 non-deterministic
	bufferedChan <- 7
	bufferedChan <- 8
	bufferedChan <- 9
	close(bufferedChan)

	<-done //chanel blocked until the data is available

	fmt.Println("main routine has completed")
}
