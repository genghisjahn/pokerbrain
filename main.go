package main

import (
	"fmt"

	"github.com/genghisjahn/pokerbrain/poker"
)

var names = []string{"Adam", "Bill", "Charles", "David", "Edward", "Frank", "Greg", "Henry", "Ivan", "Jon"}

type Combos struct {
	Hands []poker.Hand
}

func (c *Combos) Len() int           { return len(c.Hands) }
func (c *Combos) Swap(i, j int)      { c.Hands[i], c.Hands[j] = c.Hands[j], c.Hands[i] }
func (c *Combos) Less(i, j int) bool { return c.Hands[i].Score < c.Hands[j].Score }

func main() {
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

		// for _, p := range t.Players {
		// 	fmt.Println(p.Name, p.Pocket)
		// }
		t.Flop()
		for k := range t.Players {
			t.Players[k].SetBestHand(t.CommunityCards)
		}
		// fmt.Println("----")
		// fmt.Println("Flop:", t.CommunityCards)
		// fmt.Println("----")
		p := t.SortPlayerHands()
		// for _, i := range p {
		// 	fmt.Println(i)
		// }
		t.Turn()
		// fmt.Println("----")
		// fmt.Println("Turn:", t.CommunityCards)
		// fmt.Println("----")
		for k, _ := range t.Players {
			t.Players[k].SetBestHand(t.CommunityCards)
		}
		p = t.SortPlayerHands()
		// for _, i := range p {
		// 	fmt.Println(i)
		// }
		t.River()
		fmt.Println("----")
		fmt.Println("River:", t.CommunityCards)
		fmt.Println("----")
		for k, _ := range t.Players {
			t.Players[k].SetBestHand(t.CommunityCards)
		}
		p = t.SortPlayerHands()
		for _, i := range p {
			fmt.Println(i)
		}
		break
		if p[0].Hand.Name == "Flush" {
			break
		}
	}
}
