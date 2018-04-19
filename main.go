package main

func main() {
	card := deck{"Mehran", "Test", "Book", "Music", "Pen"}
	card = append(card, "Book123")
	card.print()
}
