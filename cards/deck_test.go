package main

import "testing"

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
