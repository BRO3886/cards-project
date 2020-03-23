package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//TODO: make a helper function to change a deck to a string
func (d deck) deckToString() string {
	sliceOfStrings := []string(d)
	return strings.Join(sliceOfStrings, ",")
}

//TODO: make a dave to file function

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.deckToString()), 0666)
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

// TODO: create a function to shuffle the deck

func (d deck) shuffleDeck() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
