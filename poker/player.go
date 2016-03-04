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

	return h
}
