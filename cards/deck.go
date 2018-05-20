package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
// This type inherits the same behavior of
// []string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Any variable of type "deck"
// now gets access to the
// "print" method
// In Go we never use this or self as
// the type var name, we name it with 1 or 2 letters
// of the new type name
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// fruits[startIndexIncluding:upToNotIncluding]
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// []byte represents a slice of asci keys values of a character
// in keyboard

// Type convertion, transforms a type of data to another
// []byte("Hi there!")
// greeting := "Hi there!"
// fmt.Println([]byte(greeting))
