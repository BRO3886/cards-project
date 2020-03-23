package main

func main() {
	cards := newDeck()
	cards.print()
	println("cslofc")
	println(cards.saveToFile("my_cards"))
	println("saved to file")
	loaded := newDeckFromFile("my_cars")
	println("loaded from file")
	loaded.print()
}
