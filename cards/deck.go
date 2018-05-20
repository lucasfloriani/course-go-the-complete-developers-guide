package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
// This type inherits the same behavior of
// []string
type deck []string

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
