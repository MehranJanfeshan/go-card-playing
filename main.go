package main


func main() {
	// cards := newDeck()
	// cards.saveToFile("myfile123")
	// cards := newDeckFromFile("mayfile124")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
