package main

import (
	"GitHub/Go/cards/utils"
	"errors"
	"os"
	"strings"
)

// Deck is a split of strings
type deck []string

// Return new deck type from file at filename and error passed from os.ReadFile. If filename is not given, returns a new complete deck
func newDeck(filename ...string) (d deck, e error) {
	if len(filename) > 1 {
		d = deck{}
		e = errors.New("cannot read from multiple files")
	} else if len(filename) == 1 {
		var bytes []byte
		filenameJoin := strings.Join(filename, "")
		bytes, e = os.ReadFile(filenameJoin)
		if e == nil {
			deckString := string(bytes)
			d = deck(strings.Split(deckString, ","))
		} else {
			d = deck{}
		}
	} else if len(filename) == 0 {
		d = generateDeck()
	} else {
		e = errors.New("cannot determine parameters")
	}
	return
}

// Generate a new complete deck
func generateDeck() (d deck) {
	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardRanks := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, rank := range cardRanks {
			d = append(d, rank+" of "+suit)
		}
	}
	return
}

// Returns deck d as a single string separated by the string in sep
func (d deck) getDeckString(sep string) (s string) {
	s = strings.Join(d, sep)
	return
}

// Split a deck into the dealt hand h and the remainder r. Returns: h (deck), r (deck)
func deal(d deck, handSize int) (h deck, r deck) {
	h, r = d[:handSize], d[handSize:]
	return
}

// Convert comma separated deck string to byte split and write to file
func (d deck) saveToFile(filename string) (e error) {
	e = os.WriteFile(filename, []byte(d.getDeckString(",")), 0666)
	return
}

// Shuffle a deck
func (d deck) shuffle() (s deck) {
	interfaceSlice := make([]interface{}, len(d))
	for i, card := range d {
		interfaceSlice[i] = card
	}
	randomisedInterfaceSlice := utils.RandomiseSlice(interfaceSlice)
	for _, randInterface := range randomisedInterfaceSlice {
		s = append(s, randInterface.(string))
	}
	return
}
