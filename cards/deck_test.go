package main

import (
	"os"
	"testing"
)

// needs upper-case letters
func TestNewDeck(t *testing.T) {
	deck := newDeck()

	// four suits and four cards per suit => total of 16
	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", deck[0])
	}

	if deck[len(deck)-1] != "Four of Clubs" {
		t.Errorf("Expected first card of Four of Clubs, but got %v", deck[len(deck)-1])
	}
}

func TestSaveToDeckAnNewDeckFromFile(t *testing.T) {
	preTestError := os.Remove("_decktesting")
	ignoreError := "remove _decktesting: no such file or directory"
	if preTestError != nil && preTestError.Error() != ignoreError {
		t.Errorf("Some error occurred during pre-test cleaning: %v", preTestError)
	}

	d := newDeck()

	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}

	postTestError := os.Remove("_decktesting")

	if postTestError != nil {
		t.Errorf("Some error occurred during post-test cleaning: %v", postTestError)
	}
}
