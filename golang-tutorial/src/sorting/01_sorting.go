package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{1, 35, 21, 2, 65, 72, 99}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	xs := []string{"a", "d", "c", "b", "e"}
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
