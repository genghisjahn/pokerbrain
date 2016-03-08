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
func (c *Combos) Less(i, j int) bool { return c.Hands[i].Score() < c.Hands[j].Score() }

func main() {
	var response string
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

	//fmt.Println("Pocket Cards:", t.Players[0].Pocket)

	// fb := []poker.Card{}
	// for _, c := range t.CommunityCards {
	// 	fb = append(fb, c)
	// }
	// fb = append(fb, t.Players[0].Pocket[0])
	// fb = append(fb, t.Players[0].Pocket[1])
	//
	// hands := poker.GetCardCombinations(fb)
	// cbo := Combos{Hands: hands}
	// sort.Sort(sort.Reverse(&cbo))
	// fmt.Println("Best Hand:", cbo.Hands[0], cbo.Hands[0].Name, cbo.Hands[0].Score())

	for _, p := range t.Players {
		//t.Players[k].SetBestHand(t.CommunityCards)
		fmt.Println(p.Name, p.Pocket)
	}
	fmt.Scanln(&response)
	t.Flop()
	for k := range t.Players {
		t.Players[k].SetBestHand(t.CommunityCards)
	}
	fmt.Println("----")
	fmt.Println("Flop:", t.CommunityCards)
	fmt.Println("----")
	p := t.SortPlayerHands()
	for _, i := range p {
		fmt.Println(i)
	}
	fmt.Scanln(&response)
	t.Turn()
	fmt.Println("----")
	fmt.Println("Turn:", t.CommunityCards)
	fmt.Println("----")
	//t.River()
}
