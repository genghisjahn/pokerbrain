package poker

import "sort"

type Table struct {
	CommunityCards [5]Card
	Hands          []Hand
	Players        []Player
	Deck
}

func (t *Table) Flop() {
	t.CommunityCards[0] = t.Deal()
	t.CommunityCards[1] = t.Deal()
	t.CommunityCards[2] = t.Deal()
}

func (t Table) Turn() {
	t.CommunityCards[3] = t.Deal()
}

func (t Table) River() {
	t.CommunityCards[4] = t.Deal()
}

func (t Table) CompareHands() []Hand {
	winners := []Hand{}
	sort.Sort(sort.Reverse(&t))
	for _, h := range t.Hands {
		sort.Sort(sort.Reverse(&h))
		winners = append(winners, h)
	}
	return winners
}

func (t *Table) Len() int           { return len(t.Players) }
func (t *Table) Swap(i, j int)      { t.Players[i], t.Players[j] = t.Players[j], t.Players[i] }
func (t *Table) Less(i, j int) bool { return t.Players[i].Hand.Score() < t.Players[j].Hand.Score() }
