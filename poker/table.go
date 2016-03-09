package poker

import "sort"

type Table struct {
	CommunityCards []Card   `json:"community_cards"`
	Players        []Player `json:"players"`
	Deck           `json:"deck"`
}

func (t *Table) Flop() {
	t.CommunityCards = make([]Card, 3, 5)
	t.CommunityCards[0] = t.Deal()
	t.CommunityCards[1] = t.Deal()
	t.CommunityCards[2] = t.Deal()
	for _, p := range t.Players {
		p.SetBestHand(t.CommunityCards)
	}
}

func (t *Table) Turn() {
	t.CommunityCards = append(t.CommunityCards, t.Deal())
	for _, p := range t.Players {
		p.SetBestHand(t.CommunityCards)
	}
}

func (t *Table) River() {
	t.CommunityCards = append(t.CommunityCards, t.Deal())
}

func scoreAllHands(t *Table) {
	for _, v := range t.Players {
		v.Hand.SetScore()
	}
}

func (t Table) SortPlayerHands() []Player {
	scoreAllHands(&t)
	winners := []Player{}
	sort.Sort(sort.Reverse(&t))
	for _, h := range t.Players {
		sort.Sort(sort.Reverse(&h))
		winners = append(winners, h)
	}
	return winners
}

func (t *Table) Len() int      { return len(t.Players) }
func (t *Table) Swap(i, j int) { t.Players[i], t.Players[j] = t.Players[j], t.Players[i] }
func (t *Table) Less(i, j int) bool {
	return t.Players[i].Hand.Score < t.Players[j].Hand.Score
}
