package main

import (
	"testing"

	p "github.com/genghisjahn/pokerbrain/poker"
)

var suits = []string{"♤", "♡", "♢", "♧"}

func TestScoreHandPair(t *testing.T) {
	h := p.Hand{}
	h.Cards[0] = p.Card{Name: "A", High: 14, Low: 1, Suit: "♡"}
	h.Cards[1] = p.Card{Name: "A", High: 14, Low: 1, Suit: "♧"}
	h.Cards[2] = p.Card{Name: "8", High: 8, Low: 8, Suit: "♤"}
	h.Cards[3] = p.Card{Name: "2", High: 2, Low: 2, Suit: "♤"}
	h.Cards[4] = p.Card{Name: "Q", High: 12, Low: 12, Suit: "♢"}
	h.Score()
	if h.Name != "Pair" {
		t.Errorf("Expected Pair, got %s\n", h.Name)
	}
	if h.Score() != "000000000000000141414120802" {
		t.Errorf("Expected 000000000000000141414120802, got %s\n", h.Score())
	}

	/*
		  Low  int
			High int
			Suit string
			Name string
	*/
}
