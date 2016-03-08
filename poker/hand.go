package poker

import (
	"fmt"
	"sort"
)

type Hand struct {
	Cards [5]Card
	Name  string
	Score string
}

type Pocket struct {
	Cards [2]Card
}

func (p Pocket) String() string {
	return fmt.Sprintf("%s,%s", p.Cards[0], p.Cards[1])
}

func (h *Hand) Len() int           { return len(h.Cards) }
func (h *Hand) Swap(i, j int)      { h.Cards[i], h.Cards[j] = h.Cards[j], h.Cards[i] }
func (h *Hand) Less(i, j int) bool { return h.Cards[i].High < h.Cards[j].High }

func (p *Pocket) Len() int           { return len(p.Cards) }
func (p *Pocket) Swap(i, j int)      { p.Cards[i], p.Cards[j] = p.Cards[j], p.Cards[i] }
func (p *Pocket) Less(i, j int) bool { return p.Cards[i].High < p.Cards[j].High }

func (h *Hand) SetScore() {
	var flush bool
	var straight bool
	var unique bool
	suit := h.Cards[0].Suit
	flush = true
	for _, c := range h.Cards {
		if c.Suit != suit {
			flush = false
		}
	}
	unique = checkunique(h.Cards)
	straight = checkstraight(h.Cards)
	ranks, vals := checkranks(h.Cards)
	sort.Sort(h)
	s := make([]int, 15, 15)
	s[14] = h.Cards[0].High
	s[13] = h.Cards[1].High
	s[12] = h.Cards[2].High
	s[11] = h.Cards[3].High
	s[10] = h.Cards[4].High

	if straight && flush {
		h.Name = "Straight Flush"
		s[0] = 1
		valint := getfinalscore(s)
		h.Score = valint
		return
	}
	if ranks == FullHouse {
		h.Name = "Full House"
		s[2] = vals[0]
		s[3] = vals[1]
		valint := getfinalscore(s)
		h.Score = valint
		return
	}
	if flush {
		h.Name = "Flush"
		s[4] = 1
		valint := getfinalscore(s)
		h.Score = valint
		return
	}
	if straight {
		h.Name = "Straight"
		s[5] = 1
		valint := getfinalscore(s)
		h.Score = valint
		return
	}
	if ranks != "" {
		if ranks == FourofKind {
			s[1] = vals[0]
			h.Name = "Four of a Kind"
		}
		if ranks == ThreeOfKind {
			s[6] = vals[0]
			h.Name = "Three of a Kind"
		}
		if ranks == TwoPair {
			s[7] = vals[0]
			s[8] = vals[1]
			h.Name = "Two Pair"
		}
		if ranks == Pair {
			s[9] = vals[0]
			h.Name = "Pair"
		}
		valint := getfinalscore(s)
		h.Score = valint
		return
	}

	if !straight && !flush && unique {
		h.Name = "High Card"
		valint := getfinalscore(s)
		h.Score = valint
		return
	}
	h.Score = "0"

}
