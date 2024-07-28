package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	// if (handSize > len(d)) {
	// throw an error
	// }

	hand := d[:handSize]
	remainingDeck := d[handSize:]

	return hand, remainingDeck
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// TODO: refactor with non-deprecated functions
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
