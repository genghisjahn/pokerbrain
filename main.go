package main

import (
	"fmt"

	"github.com/genghisjahn/pokerbrain/poker"
)

var names = []string{"Adam", "Bill", "Charles", "David", "Edward", "Frank", "Greg", "Henry", "Ivan", "Jon"}

func main() {
	t := poker.Table{}
	deck := poker.BuildDeck()
	deck.Shuffle(5)
	t.Deck = deck
	players := make([]poker.Player, 10, 10)

	for k := range players {
		players[k].Name = names[k]
		players[k].Position = k + 1
		players[k].Stake = 100.0
		players[k].Pocket[0] = t.Deck.Deal()
		players[k].Pocket[1] = t.Deck.Deal()
	}
	t.Players = append(t.Players, players...)
	t.Flop()

	for k, p := range t.Players {
		t.Players[k].SetBestHand(t.CommunityCards)
		fmt.Println(p.Name, p.Pocket)
	}
	fmt.Println(t.CommunityCards)
	fmt.Println("----")
	p := t.SortPlayerHands()
	for _, i := range p {
		fmt.Println(i)
	}
}
