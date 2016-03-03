package poker

import "sort"

type Table struct {
	CommunityCards [5]Card
	Hands          []Hand
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

func (t *Table) Len() int           { return len(t.Hands) }
func (t *Table) Swap(i, j int)      { t.Hands[i], t.Hands[j] = t.Hands[j], t.Hands[i] }
func (t *Table) Less(i, j int) bool { return t.Hands[i].Score() < t.Hands[j].Score() }
