package main

import (
	"github.com/cozmus/channels/helpers"
	"log"
)

const numPool = 1000

func CalculateValue(someChan chan int){
	randomNumber := helpers.GenerateRandomNumber(numPool)
	someChan <- randomNumber //passes value into the channel
}

//enable Go modules integration from GoLand preferences
func main()  {
	intChan := make(chan int)
	defer close(intChan) //good practise to close channels

	go CalculateValue(intChan) //creates new go routine

	num := <- intChan //listening value from channel and assigned when done
	log.Println(num)
}


