package main

import (
	"fmt"
	"time"
)

func main() {
	unbufferedChan := make(chan int)

	//reader go routine
	go func(chanelName chan int) {
		val := <-chanelName
		fmt.Println(val)
	}(unbufferedChan)

	//writer go routine
	go func(chanelName chan int) {
		chanelName <- 11
	}(unbufferedChan)

	time.Sleep(time.Second)
	fmt.Println("Done")
}
