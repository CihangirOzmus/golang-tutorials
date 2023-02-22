package main

import (
	"os"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	expectedLenOfDeck := 20
	d := createDeck()

	if len(d) != expectedLenOfDeck {
		t.Errorf("Expected deck length of %d but got %d", expectedLenOfDeck, len(d))
	}

	expectedFirstCard := "Ace of Hearts"
	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card of the deck is %s, but got %s", expectedFirstCard, d[0])
	}

	expectedLastCard := "Five of Spades"
	lastIdx := len(d) - 1
	if d[lastIdx] != expectedLastCard {
		t.Errorf("Expected last card of the deck is %s, but got %s", expectedLastCard, d[lastIdx])
	}
}

func TestSaveToFile(t *testing.T) {
	testingFile := "_deckTesting"
	os.Remove(testingFile)

	d := createDeck()
	d.saveToFile(testingFile)

	loadedDeck := readFromFile(testingFile)
	expectedLenOfDeck := 20

	if len(loadedDeck) != expectedLenOfDeck {
		t.Errorf("Expected deck length of %d but got %d", expectedLenOfDeck, len(loadedDeck))
	}

	os.Remove(testingFile)
}
