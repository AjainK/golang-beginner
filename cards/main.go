package main

func main() {
	// cards := newDeck()
	// fmt.Println(len(cards))
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// fmt.Println(len(cards))
	cards := newDeck()
	cards.print()

	//cards := newDeckFromFile("myCards")
	cards.shuffle()
	cards.saveToFile("myCards")
	cards.print()
}
