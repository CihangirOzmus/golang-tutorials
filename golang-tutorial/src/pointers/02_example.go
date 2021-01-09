package main

import "fmt"

type person struct {
	first string
}

func main() {
	p := person{"james"}
	fmt.Println(p.first)
	changeMe(&p)
	fmt.Println(p.first)
}

func changeMe(p *person) {
	p.first = "jonny"
	(*p).first = "jane"
}
