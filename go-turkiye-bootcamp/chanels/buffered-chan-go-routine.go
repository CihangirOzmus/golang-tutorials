package main

import "fmt"

func main() {
	bufferedChan := make(chan int, 5)

	go func(chanelName chan int) {
		for {
			val := <-chanelName
			fmt.Println(val)
		}
	}(bufferedChan)

	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3
	bufferedChan <- 4
	bufferedChan <- 5
	bufferedChan <- 6 // after 5 non-deterministic
	bufferedChan <- 7
	bufferedChan <- 8
	bufferedChan <- 9
}
