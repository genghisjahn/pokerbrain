package poker

type Player struct {
	Name string
	Hand
	Position int
	Stake    float64
	Pocket   [2]Card
}

func (p *Player) SetBestHand(cc []Card) {
	h := Hand{}
	//Right now, just test a flop, 3 community cards
	h.Cards[0] = p.Pocket[0]
	h.Cards[1] = p.Pocket[1]
	h.Cards[2] = cc[0]
	h.Cards[3] = cc[1]
	h.Cards[4] = cc[2]
	p.Hand = h
}
