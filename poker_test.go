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
	h.SetScore()
	if h.Name != "Pair" {
		t.Errorf("Expected Pair, got %s\n", h.Name)
	}
	if h.Score != "000000000000000141414120802" {
		t.Errorf("Expected 000000000000000141414120802, got %s\n", h.Score)
	}
}

func TestScoreHandTwoPair(t *testing.T) {
	h := p.Hand{}
	h.Cards[0] = p.Card{Name: "A", High: 14, Low: 1, Suit: "♡"}
	h.Cards[1] = p.Card{Name: "A", High: 14, Low: 1, Suit: "♧"}
	h.Cards[2] = p.Card{Name: "8", High: 8, Low: 8, Suit: "♤"}
	h.Cards[3] = p.Card{Name: "2", High: 2, Low: 2, Suit: "♤"}
	h.Cards[4] = p.Card{Name: "8", High: 8, Low: 8, Suit: "♢"}
	h.SetScore()
	if h.Name != "Two Pair" {
		t.Errorf("Expected Two Pair, got %s\n", h.Name)
	}
	if h.Score != "000000000001408001414080802" {
		t.Errorf("Expected 000000000001408001414080802, got %s\n", h.Score)
	}
}

func TestScoreHandThree(t *testing.T) {
	h := p.Hand{}
	h.Cards[0] = p.Card{Name: "A", High: 14, Low: 1, Suit: "♡"}
	h.Cards[1] = p.Card{Name: "A", High: 14, Low: 1, Suit: "♧"}
	h.Cards[2] = p.Card{Name: "A", High: 14, Low: 1, Suit: "♤"}
	h.Cards[3] = p.Card{Name: "2", High: 2, Low: 2, Suit: "♤"}
	h.Cards[4] = p.Card{Name: "8", High: 8, Low: 8, Suit: "♢"}
	h.SetScore()
	if h.Name != "Three of a Kind" {
		t.Errorf("Expected Three of a Kind, got %s\n", h.Name)
	}
	if h.Score != "000000000140000001414140802" {
		t.Errorf("Expected 000000000140000001414140802, got %s\n", h.Score)
	}
}

func TestScoreHandStraight(t *testing.T) {
	h := p.Hand{}
	h.Cards[0] = p.Card{Name: "2", High: 2, Low: 2, Suit: "♡"}
	h.Cards[1] = p.Card{Name: "3", High: 3, Low: 3, Suit: "♧"}
	h.Cards[2] = p.Card{Name: "6", High: 6, Low: 6, Suit: "♤"}
	h.Cards[3] = p.Card{Name: "4", High: 4, Low: 4, Suit: "♤"}
	h.Cards[4] = p.Card{Name: "5", High: 5, Low: 5, Suit: "♢"}
	h.SetScore()
	if h.Name != "Straight" {
		t.Errorf("Expected Straight, got %s\n", h.Name)
	}
	if h.Score != "000000001000000000605040302" {
		t.Errorf("Expected 000000001000000000605040302, got %s\n", h.Score)
	}
}
