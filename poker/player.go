package poker

import "sort"

type Player struct {
	Name string
	Hand
	Position int
	Stake    float64
	Pocket   [2]Card
}

type Combos struct {
	Hands []Hand
}

func (c *Combos) Len() int           { return len(c.Hands) }
func (c *Combos) Swap(i, j int)      { c.Hands[i], c.Hands[j] = c.Hands[j], c.Hands[i] }
func (c *Combos) Less(i, j int) bool { return c.Hands[i].Score < c.Hands[j].Score }

func (p *Player) SetBestHand(cc []Card) {
	fb := []Card{}
	for _, c := range cc {
		fb = append(fb, c)
	}
	fb = append(fb, p.Pocket[0])
	fb = append(fb, p.Pocket[1])
	hands := GetCardCombinations(fb)
	cbo := Combos{Hands: hands}
	sort.Sort(sort.Reverse(&cbo)) //GetCardCombinations(cards []Card)
	p.Hand = cbo.Hands[0]

}
