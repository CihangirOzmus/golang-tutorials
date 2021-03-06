package main

import "log"

var s = "seven"

func main()  {
	var s2 = "six"

	s := "eight" //give attention to scope

	log.Println("s is", s) //uses block scope
	log.Println("s2 is", s2)
	saySometing("XXX")
}

func saySometing(s3 string) (string, string)  {
	log.Println("s from func call", s) //uses global scope
	return s3, "World"
}