package main


func main() {
	// var card string = "Ace of Spades"
	// cards := newDeck()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards.print()

	cards := newDeckFromFile("my_cards")
	cards.print()
	
}


