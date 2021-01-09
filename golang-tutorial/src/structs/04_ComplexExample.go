package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle:   vehicle{doors: 4, color: "red"},
		fourWheel: true,
	}

	s := sedan{
		vehicle{4, "black"},
		true,
	}

	fmt.Println(t)
	fmt.Println(s)

}
