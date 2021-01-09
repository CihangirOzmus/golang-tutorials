package main

import "fmt"

//official doc says use slices not arrays
func main() {
	var fixSizeNumbers = [3]int{10, 20, 30}
	fmt.Println(fixSizeNumbers)
	fmt.Println("Size of array: ", len(fixSizeNumbers))

	var numbers = [...]int{10, 20, 30, 40}
	fmt.Println(numbers)

	var emptyArray [5]string
	fmt.Println(emptyArray)

	emptyArray[2] = "Hi"
	fmt.Println(emptyArray)
}
