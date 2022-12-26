package main

import (
	"fmt"
)

func main() {
	bufferedChan := make(chan int, 1)
	bufferedChan <- 1

	select { //listen chanel with select case
	case val := <-bufferedChan:
		fmt.Println(val)
	}
}
