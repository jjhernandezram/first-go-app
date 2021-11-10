package main

func main() {
	//cards := []string{newCard(), "Ace of Spades"}
	cards := newDeck()
	//hand, remainingCards := deal(cards, 5)
	//
	//fmt.Println("deck: ", len(cards))
	//fmt.Println("Hand: ", len(hand))
	//fmt.Println("Remaining: ", len(remainingCards))
	//
	//cards.saveToFile("deck_cards")
	//
	//cards2 := newDeckFromFile("deck_cards")
	//fmt.Println(len(cards2))

	cards.shuffle()
	cards.print()
}
