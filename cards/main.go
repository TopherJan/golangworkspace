package main

import (
	"fmt"
	"os"
)

func main() {
	// Create a deck
	deck := newDeck()

	// Print out the values in the deck
	fmt.Println("----DECK-----")
	deck.print()

	// Deal 10 cards the deck
	hand, remainingDeck := deal(deck, 10)

	// Print out the values of the cards in your hand
	fmt.Println("----HAND-----")
	hand.print()

	// Print out the values of the remaining dock
	fmt.Println("----REMAINING DECK-----")
	remainingDeck.print()

	// Save the remaining deck into a file
	//
	err := remainingDeck.saveToFile("my_cards")
	if err != nil {
		fmt.Println("Failed to save into file with error:", err.Error())
		os.Exit(-1)
	}

	// Create a new deck from a file
	newDeck := newDeckFromFile("my_cards")

	// Print out the values in the new deck
	fmt.Println("----DECK FROM FILE-----")
	newDeck.print()

	// Print out the values in the new deck
	fmt.Println("----SHUFFLED DECK FROM FILE-----")
	newDeck.shuffle()
	newDeck.print()
}
