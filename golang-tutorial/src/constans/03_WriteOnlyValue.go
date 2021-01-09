package main

import "fmt"

const (
	_ = iota //ignore first value by assigning to blank identifier
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

func main() {
	var specialistType int //takes zero iota
	fmt.Println(specialistType == catSpecialist)
}
