package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type shape interface {
	area() float64
}

func main() {
	c := circle{3}
	s := square{3}

	info(c)
	info(s)
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Phi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println(s.area())
}
