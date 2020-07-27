package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck len of 16 but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Club" {
		t.Errorf("Expected Four of Club but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_testFile")
	deck := newDeck()
	deck.saveToFile("_testFile")
	loadedDeck := newDeckFromFile("_testFile")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected len 16 got %v", len(loadedDeck))
	}
	os.Remove("_testFile")
}
