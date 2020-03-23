package main

func main() {
	cards := newDeck()
	cards.print()
	println("cslofc")
	println(cards.saveToFile("my_cards"))

}
