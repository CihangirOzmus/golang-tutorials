package main

import (
	"fmt"
)

const myConst int16 = 27

func main() {
	//uppercase means global, so dont use if you dont need
	fmt.Printf("%v. %T\n", myConst, myConst)

	const myConst int = 42
	fmt.Printf("%v. %T\n", myConst, myConst)
}
