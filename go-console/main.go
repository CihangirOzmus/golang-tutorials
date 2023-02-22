package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
	"strconv"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappuccino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit")

	ch := ' '
	for ch != 'q' && ch != 'Q' {
		ch, _, err = keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		coffeeMenu := make(map[int]string)
		coffeeMenu[1] = "Cappuccino"
		coffeeMenu[2] = "Latte"
		coffeeMenu[3] = "Americano"
		coffeeMenu[4] = "Mocha"
		coffeeMenu[5] = "Macchiato"
		coffeeMenu[6] = "Espresso"

		entryNum, _ := strconv.Atoi(string(ch))
		_, ok := coffeeMenu[entryNum]
		if ok {
			fmt.Println(fmt.Sprintf("%s is chosen", coffeeMenu[entryNum]))
		}
	}

	fmt.Println("Program exiting!")
}
