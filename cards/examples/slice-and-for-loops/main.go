package main

import "fmt"

func main() {
	// Create a slice
	cards := []string{"Ace of Diamonds", newCard()}
	// Add new value to slice an return the result to add in the same variable
	cards = append(cards, "Six of Spades")

	// User range to iterate all values from a slice/array.
	// range returns a key => value each time of iteration
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
