package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// TODO: Create a new type of deck
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"ace", "two", "three", "four", "five"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (cards deck) print() {
	for _, card := range cards {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//TODO: make a helper function to change a deck to a string
func (cards deck) deckToString() string {
	sliceOfStrings := []string(cards)
	return strings.Join(sliceOfStrings, ",")
}

//TODO: make a dave to file function

func (cards deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(cards.deckToString()), 0666)
}

//TODO: make a function to read form file

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	// if something went wrong during the process, error has some value. If everything is right, error is nil
	if err != nil {
		//option1 - log the err and return new deck using newDeck function
		//option1 - log the err and quit the program
		fmt.Println("error:", err)
		os.Exit(1)
	}
	d := strings.Split(string(byteSlice), ",")
	return deck(d)
}
