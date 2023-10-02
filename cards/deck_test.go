package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Check the length of the deck
	//
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d.", len(d))
	}

	// Check the first card in the deck
	//
	firstCard := d[0]
	if firstCard != "Ace of Spades" {
		t.Errorf("Expected first card of 'Ace of Spades' but got %s", firstCard)
	}

	// Check the last card in the deck
	//
	lastCard := d[len(d)-1]
	if lastCard != "King of Clubs" {
		t.Errorf("Expected first card of 'King of Clubs' but got %s", lastCard)
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	err := deck.saveToFile(filename)
	if err != nil {
		t.Errorf("Failed to save deck into file %s", filename)
	}

	loadedDeck := newDeckFromFile(filename)

	// Check the length of the deck
	//
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d.", len(loadedDeck))
	}

	os.Remove(filename)
}
