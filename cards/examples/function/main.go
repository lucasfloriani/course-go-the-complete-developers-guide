package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

func newCard() /*int*/ string {
	return "Five of Diamonds"
	// return 12
}
