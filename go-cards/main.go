package main

func main() {
	cards := createDeck()
	cards.shuffle()
	cards.printDeck()

	cardsFromFile := readFromFile("my_cards")
	cardsFromFile.printDeck()
}
