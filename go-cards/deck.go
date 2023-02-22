package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func createDeck() deck {
	cards := deck{}
	cardShapes := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardNumbers := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, shape := range cardShapes {
		for _, num := range cardNumbers {
			cards = append(cards, num+" of "+shape)
		}
	}

	return cards
}

func (d deck) printDeck() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) dealDeck(handSize int) (deck, deck, error) {
	if handSize >= len(d) {
		return nil, nil, errors.New("hand size should be lower than the deck size")
	}
	return d[:handSize], d[handSize:], nil
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		rIdx := r.Intn(len(d) - 1)
		d[i], d[rIdx] = d[rIdx], d[i]
	}
}

func (d deck) saveToFile(filename string) {
	stringDeck := strings.Join(d, ",")
	err := os.WriteFile(filename, []byte(stringDeck), 0666)
	if err != nil {
		log.Fatal("Can not write to file.")
	}
}

func readFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Can not read from file.")
	}

	s := strings.Split(string(bs), ",")
	return s
}
