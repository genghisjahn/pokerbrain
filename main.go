package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"

	"github.com/genghisjahn/pokerbrain/poker"
)

var suits = []string{"♤", "♡", "♢", "♧"}

const (
	HighCard      = "High Card"
	Pair          = "Pair"
	TwoPair       = "Two Pair"
	ThreeOfKind   = "Three of a Kind"
	Straight      = "Straight"
	Flush         = "Flush"
	FullHouse     = "Full House"
	FourofKind    = "Four of a Kind"
	StraightFlush = "Straight Flush"
	LOW           = false
	HIGH          = true
)

type deck struct {
	Cards []poker.Card
}

type player struct {
	Cards [2]poker.Card
}

type table struct {
	Cards [5]poker.Card
	Hands []hand
}

type hand struct {
	Cards [5]poker.Card
	Name  string
}

func main() {
	deck := buildDeck()
	deck.Shuffle(5)
	hands := make([]hand, 10, 10)

	deck = buildDeck()
	deck.Shuffle(5)
	for i, h := range hands {
		for k := range h.Cards {
			hands[i].Cards[k] = deck.Deal()
		}
	}

	t := table{}
	t.Hands = append(t.Hands, hands...)

	ch := compareHands(t)
	fmt.Println("Compare:")
	for _, c := range ch {
		fmt.Println(c)
	}

}

func (d *deck) Deal() poker.Card {
	card := d.Cards[0]
	d.Cards = append(d.Cards[:0], d.Cards[1:]...)
	return card
}

func getfinalscore(vals []int) string {
	strval := ""
	pz := map[int]int{1: 1, 2: 2, 3: 3, 6: 6, 7: 7, 8: 8, 9: 9, 10: 10, 11: 11, 12: 12, 13: 13, 14: 14}
	for k, v := range vals {
		a := strconv.Itoa(v)
		if _, ok := pz[k]; ok {

			if v < 10 {
				a = "0" + a
			}
		}
		strval += a
	}
	return strval
}

func (h *hand) Score() string {
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
		return valint
	}
	if ranks == FullHouse {
		h.Name = "Full House"
		s[2] = vals[0]
		s[3] = vals[1]
		valint := getfinalscore(s)
		return valint
	}
	if flush {
		h.Name = "Flush"
		s[4] = 1
		valint := getfinalscore(s)
		return valint
	}
	if straight {
		h.Name = "Straight"
		s[5] = 1
		valint := getfinalscore(s)
		return valint
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
		return valint
	}

	if !straight && !flush && unique {
		h.Name = "High Card"
		valint := getfinalscore(s)
		return valint
	}
	return "0"
}

type rankResult struct {
	Values [15]byte
}

func checkranks(cards [5]poker.Card) (string, [2]int) {
	c := make(map[int]int)
	for _, v := range cards {
		if _, ok := c[v.High]; ok {
			c[v.High]++
			continue
		}
		c[v.High] = 1
	}

	var kickers = []poker.Card{}
	for _, v := range cards {
		if c[v.High] == 1 {
			kickers = append(kickers, v)
		}
	}

	var onepair bool
	var twopair bool
	var threekind bool
	var fourkind bool
	for _, v := range c {
		if v == 2 {
			if onepair {
				twopair = true
			}
			onepair = true
		}
		if v == 3 {
			threekind = true
		}
		if v == 4 {
			fourkind = true
		}
	}

	vals := [2]int{}
	if fourkind {
		for k, c1 := range c {
			if c1 == 4 {
				vals[0] = k
			}
		}
		return FourofKind, vals
	}
	if threekind && onepair {
		for k, c1 := range c {
			if c1 == 3 {
				vals[0] = k
			}
			if c1 == 2 {
				vals[1] = k
			}
		}
		return FullHouse, vals
	}
	if threekind && !onepair {
		for k, c1 := range c {
			if c1 == 3 {
				vals[0] = k
			}
		}
		return ThreeOfKind, vals
	}
	if twopair {
		for k, c1 := range c {
			if c1 == 2 {
				vals[0] = k
			}
			if c1 == 2 {
				vals[1] = k
			}
		}
		if vals[1] > vals[0] {
			vals[0], vals[1] = vals[1], vals[0]
		}
		return TwoPair, vals
	}
	if onepair {
		for k, c1 := range c {
			if c1 == 2 {
				vals[0] = k
			}
		}
		return Pair, vals
	}
	return "", vals
}

func checkunique(cards [5]poker.Card) bool {
	c := make(map[int]int)
	for _, v := range cards {
		if _, ok := c[v.High]; ok {
			return false
		}
		c[v.High] = v.High
	}
	return true
}

func checkstraight(cards [5]poker.Card) bool {
	//replace this with checkunique
	if !checkunique(cards) {
		return false
	}

	highMin, highMax := getbounds(HIGH, cards)
	lowMin, lowMax := getbounds(LOW, cards)
	if highMax-highMin == 4 {
		return true
	}
	if lowMax-lowMin == 4 {
		return true
	}
	return false
}

func checkMaxDiff(min, max int) bool {
	if max-min == 5 {
		return true
	}
	return false
}

func getbounds(highlow bool, c [5]poker.Card) (int, int) {
	min := 15
	max := 0
	vals := [5]int{}
	for k := range c {
		if highlow == HIGH {
			vals[k] = c[k].High
			continue
		}
		vals[k] = c[k].Low
	}
	for _, v := range vals {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func (d *deck) Shuffle(num int) {
	rand.Seed(time.Now().UnixNano())
	for n := 0; n < num; n++ {
		for i := range d.Cards {
			j := rand.Intn(i + 1)
			d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
		}
	}
}

func compareHands(t table) []hand {
	winners := []hand{}
	sort.Sort(sort.Reverse(&t))
	for _, h := range t.Hands {
		sort.Sort(sort.Reverse(&h))
		winners = append(winners, h)
	}
	return winners
}

func (h *hand) Len() int           { return len(h.Cards) }
func (h *hand) Swap(i, j int)      { h.Cards[i], h.Cards[j] = h.Cards[j], h.Cards[i] }
func (h *hand) Less(i, j int) bool { return h.Cards[i].High < h.Cards[j].High }

func (t *table) Len() int           { return len(t.Hands) }
func (t *table) Swap(i, j int)      { t.Hands[i], t.Hands[j] = t.Hands[j], t.Hands[i] }
func (t *table) Less(i, j int) bool { return t.Hands[i].Score() < t.Hands[j].Score() }

func buildDeck() deck {
	var d = deck{}
	for _, v := range suits {
		for i := 1; i < 14; i++ {

			c := poker.Card{Low: i, Suit: v, High: i}
			if i == 1 {
				c.High = 14
				c.Name = "A"
			}
			if i > 1 && i < 11 {
				c.Name = fmt.Sprintf("%v", i)
			}

			if i == 11 {
				c.Name = "J"
			}
			if i == 12 {
				c.Name = "Q"
			}
			if i == 13 {
				c.Name = "K"
			}
			d.Cards = append(d.Cards, c)
		}
	}
	return d
}
