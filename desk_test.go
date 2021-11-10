package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 56 {
		t.Errorf("Expedted deck size of 56, but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expedted 'Ace of Spades' but got %s", d[0])
	}

	if d[(len(d)-1)] != "King of Clubs" {
		t.Errorf("Expedted 'King of Clubs' but got %s", d[(len(d)-1)])
	}
}

func TestSaveToFileNewDeckFromFile(t *testing.T) {
	os.Remove("_testdeckfile")

	deck := newDeck()
	deck.saveToFile("_testdeckfile")

	loadedDeck := newDeckFromFile("_testdeckfile")

	if len(loadedDeck) != 56 {
		t.Errorf("Expedted deck size of 56, but got %d", len(loadedDeck))
	}

	os.Remove("_testdeckfile")
}
