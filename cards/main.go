package main

import (
	"fmt"
)

func main() {
	cards, _ := newDeck()
	cards.saveToFile("cards.txt")

	fmt.Println(cards.getDeckString(","))
	cardsShuffle := cards.shuffle()
	fmt.Println(cardsShuffle.getDeckString(","))
}
