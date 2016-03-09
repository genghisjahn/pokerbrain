package demo

import (
	"fmt"

	"github.com/genghisjahn/pokerbrain/poker"
)

var names = []string{"Adam", "Bill", "Charles", "David", "Edward", "Frank", "Greg", "Henry", "Ivan", "Jon"}

func Demotable() {
	for {

		t := poker.Table{}
		deck := poker.BuildDeck()
		deck.Shuffle(5)
		t.Deck = deck
		players := make([]poker.Player, 0, 10)

		for k := range players[:10] {
			newPlayer := poker.Player{}
			newPlayer.Name = names[k]
			newPlayer.Position = k + 1
			newPlayer.Stake = 100.0
			newPlayer.Pocket[0] = t.Deck.Deal()
			newPlayer.Pocket[1] = t.Deck.Deal()
			players = append(players, newPlayer)
		}
		t.Players = append(t.Players, players...)
		t.Flop()
		for k := range t.Players {
			t.Players[k].SetBestHand(t.CommunityCards)
		}
		p := t.SortPlayerHands()
		t.Turn()
		for k := range t.Players {
			t.Players[k].SetBestHand(t.CommunityCards)
		}
		p = t.SortPlayerHands()
		t.River()
		fmt.Println("----")
		fmt.Println("River:", t.CommunityCards)
		fmt.Println("----")
		for k := range t.Players {
			t.Players[k].SetBestHand(t.CommunityCards)
		}
		p = t.SortPlayerHands()
		for _, i := range p {
			fmt.Println(i)
		}
		if p[0].Hand.Name == "Pair" {
			break
		}
	}
}
