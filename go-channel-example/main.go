package main

import (
	"bufio"
	"fmt"
	"go-channel-example/logger"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)

	go logger.ListenForLog(ch)

	for i := 0; i < 5; i++ {
		fmt.Println("-> ")
		userInput, _ := reader.ReadString('\n')
		ch <- userInput
		time.Sleep(time.Second)
	}
}
