package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	cards.print()
}

func newCard() /*int*/ string {
	return "Five of Diamonds"
	// return 12
}
