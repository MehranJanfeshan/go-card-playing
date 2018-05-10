package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 12 {
		t.Errorf("Expected lenght of 12, but got %v", len(d))
	}
}

func TestSaveFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 12 {
		t.Errorf("Expected lenght of 12 but got %v" , len(loadedDeck))
	}

	os.Remove("_decktesting")
}
