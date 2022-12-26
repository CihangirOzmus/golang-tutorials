package main

import (
	"fmt"
	"time"
)

func main() {
	//var unbufferedChan chan int (nil chanels)
	unbufferedChan := make(chan int) // every read write is blocking

	go func(chanelName <-chan int) {
		val := <-chanelName
		fmt.Println(val)
	}(unbufferedChan)

	unbufferedChan <- 1
	time.Sleep(time.Second)
}
