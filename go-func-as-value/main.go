package main

import (
	"fmt"
	"log"
	"strconv"
)

type opFuncType func(int, int) int

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
	}

	operations := map[string]opFuncType{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	for _, expression := range expressions {
		if len(expressions) < 3 {
			log.Println("Invalid expression:", expression)
			continue
		}

		fNum, _ := strconv.Atoi(expression[0])
		op := expression[1]
		opFunc, ok := operations[op]
		if !ok {
			log.Println("unsupported operation!", op)
			continue
		}
		sNum, _ := strconv.Atoi(expression[2])

		fmt.Println(opFunc(fNum, sNum))
	}
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	if b == 0 {
		log.Fatal("Divider should not be zero!")
	}
	return a / b
}
