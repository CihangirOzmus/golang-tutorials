package main

import "log"

func main()  {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)

	changeUsingPointer(&myString) //reference to memory

	log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string)  { //passes pointer to variable
	log.Println("s adress: ", s)
	newValue := "Red"
	*s = newValue
}