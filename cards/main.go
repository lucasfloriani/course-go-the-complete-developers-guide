package main

import "fmt"

func main() {
	// // var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds" // Dont use column to reasign because the var already know it is for string values

	card := newCard()

	fmt.Println(card)
}

func newCard() /*int*/ string {
	return "Five of Diamonds"
	// return 12
}
