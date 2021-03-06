package main

import (
	"log"
	"sort"
)

//use slice instead of array
//slice keeps order
func main()  {
	var mySlice []string

	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Mary")
	log.Println(mySlice)

	var mySlice2 []int
	mySlice2 = append(mySlice2, 2)
	mySlice2 = append(mySlice2, 1)
	mySlice2 = append(mySlice2, 3)

	sort.Ints(mySlice2)
	log.Println(mySlice2)

	numbers := []int{1,2,3,4,5,6,7,8,9,10}
	log.Println(numbers)
	log.Println(numbers[0:3])
	log.Println(numbers[6:9])

	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}