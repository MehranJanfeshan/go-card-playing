package main


func main() {
	// cards := newDeck()
	// cards.saveToFile("myfile123")
	cards := newDeckFromFile("mayfile123")
	cards.print()
}
