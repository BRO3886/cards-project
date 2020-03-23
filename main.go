package main

func main() {
	cards := newDeck()
	cards.print()
	println("cslofc")
	println(cards.saveToFile("my_cards"))
	println("\nsaved to file\n")
	loaded := newDeckFromFile("my_cards")
	println("\nloaded from file\n")
	loaded.print()
	cards.shuffleDeck()
	println("\nshuffled\n")
	cards.print()
}
