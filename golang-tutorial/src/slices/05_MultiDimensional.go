package main

import "fmt"

func main() {
	firstSlice := []string{"James", "Bond"}
	secondSlice := []string{"Strawberry", "Apple"}

	multiSlice := [][]string{firstSlice, secondSlice}
	fmt.Println(multiSlice)

	multiSlice = append(multiSlice, []string{"Hi", "there"})
	fmt.Println(multiSlice)
}
