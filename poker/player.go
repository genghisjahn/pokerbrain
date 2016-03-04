package poker

type Player struct {
	Name string
	Hand
	Position int
	Stake    float64
	Pocket   [2]Card
}

func (p *Player) BestHand(cc []Card) Hand {
	h := Hand{}
	//Right now, just test a flop, 3 community cards
	p.Hand.Cards[0] = p.Pocket[0]
	p.Hand.Cards[1] = p.Pocket[1]
	p.Hand.Cards[2] = cc[0]
	p.Hand.Cards[3] = cc[1]
	p.Hand.Cards[4] = cc[2]
	return h
}
