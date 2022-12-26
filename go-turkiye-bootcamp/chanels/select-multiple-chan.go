package main

import "fmt"

func main() {
	chan1 := make(chan int, 1)
	chan1 <- 1

	chan2 := make(chan int, 1)
	chan2 <- 2

	//select { //will print one of them, because go will select one of them
	//case val1 := <-chan1:
	//	fmt.Println(val1)
	//case val2 := <-chan2:
	//	fmt.Println(val2)
	//}

	//for { // wont work because of deadlock, channels wont get data after two loop
	//	select {
	//	case val1 := <-chan1:
	//		fmt.Println(val1)
	//	case val2 := <-chan2:
	//		fmt.Println(val2)
	//	default:
	//		break //wont work, because break will break select not for
	//	}
	//}

	var done bool
	for !done { //this works
		select {
		case val1 := <-chan1:
			fmt.Println(val1)
		case val2 := <-chan2:
			fmt.Println(val2)
		default:
			done = true
		}
	}

}
