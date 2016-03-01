package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var suits = []string{"♤", "♡", "♢", "♧"}

const (
	HighCard      = 20000
	Pair          = 40000
	TwoPair       = 60000
	ThreeOfKind   = 70000
	Straight      = 80000
	Flush         = 90000
	FullHouse     = 100000
	FourofKind    = 120000
	StraightFlush = 140000
	LOW           = false
	HIGH          = true
)

type card struct {
	Low   int
	High  int
	Suit  string
	Name  string
	Value int
}

type deck struct {
	Cards []card
}

type player struct {
	Cards [2]card
}

type table struct {
	Cards [5]card
	Hands []hand
}

type hand struct {
	Cards [5]card
	Value int
	Name  string
}

func main() {
	deck := buildDeck()
	deck.Shuffle(5)
	for h := 0; h < 100000; h++ {
		hd1 := hand{}
		hd2 := hand{}
		deck = buildDeck()
		deck.Shuffle(5)
		for k := range hd1.Cards {
			hd1.Cards[k] = deck.Deal()
			hd2.Cards[k] = deck.Deal()
		}
		hd1.Score()
		hd2.Score()
		if hd1.Name == "Pair" && hd2.Name == "Pair" {

			sort.Sort(sort.Reverse(&hd1))
			sort.Sort(sort.Reverse(&hd2))
			fmt.Println(hd1.Score(), hd1.Cards)
			fmt.Println(hd2.Score(), hd2.Cards)
			fmt.Println("-----")
			t := table{}
			t.Hands = append(t.Hands, hd1, hd2)

			ch := compareHands(t)
			fmt.Println("Compare:")
			fmt.Println(ch)
			break
		}
	}
}

func (d *deck) Deal() card {
	card := d.Cards[0]
	d.Cards = append(d.Cards[:0], d.Cards[1:]...)
	return card
}

func (h *hand) Score() int {
	var flush bool
	var straight bool
	var unique bool
	var sval int
	suit := h.Cards[0].Suit
	flush = true
	for _, c := range h.Cards {
		sval += c.Value
		if c.Suit != suit {
			flush = false
		}
	}
	unique = checkunique(h.Cards)
	straight = checkstraight(h.Cards)
	ranks := checkranks(h.Cards)
	if straight && flush {
		h.Name = "Straight Flush"
		h.Value = sval + StraightFlush
		return StraightFlush
	}
	if ranks == FullHouse {
		h.Name = "Full House"
		h.Value = sval + FullHouse
		return FullHouse
	}
	if flush {
		h.Name = "Flush"
		h.Value = sval + Flush
		return Flush
	}
	if straight {
		h.Name = "Straight"
		h.Value = sval + Straight
		return Straight
	}
	if ranks > 0 {
		if ranks == FourofKind {
			h.Name = "Four of a Kind"
		}
		if ranks == ThreeOfKind {
			h.Name = "Three of a Kind"
		}
		if ranks == TwoPair {
			h.Name = "Two Pair"
		}
		if ranks == Pair {
			h.Name = "Pair"
		}
		h.Value = sval + ranks
		return ranks
	}

	if !straight && !flush && unique {
		h.Name = "High Card"
		return HighCard + sval
	}
	return 0
}

func checkranks(cards [5]card) int {
	c := make(map[int]int)
	for _, v := range cards {
		if _, ok := c[v.High]; ok {
			c[v.High]++
			continue
		}
		c[v.High] = 1
	}

	var kickers = []card{}
	var rankCards = []card{}
	for _, v := range cards {
		if c[v.High] == 1 {
			kickers = append(kickers, v)
		} else {
			rankCards = append(rankCards, v)
		}
	}

	fmt.Println("Kickers:", kickers)

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
	if fourkind {
		return FourofKind
	}
	if threekind && onepair {
		return FullHouse
	}
	if threekind && !onepair {
		return ThreeOfKind
	}
	if twopair {
		return TwoPair
	}
	if onepair {
		return Pair
	}
	return 0
}

func checkunique(cards [5]card) bool {
	c := make(map[int]int)
	for _, v := range cards {
		if _, ok := c[v.High]; ok {
			return false
		}
		c[v.High] = v.High
	}
	return true
}

func checkstraight(cards [5]card) bool {
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

func getbounds(highlow bool, c [5]card) (int, int) {
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

func (c card) String() string {
	return fmt.Sprintf("%s%s", c.Name, c.Suit)
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
		winners = append(winners, h)
	}
	return winners
}

func (h *hand) Len() int           { return len(h.Cards) }
func (h *hand) Swap(i, j int)      { h.Cards[i], h.Cards[j] = h.Cards[j], h.Cards[i] }
func (h *hand) Less(i, j int) bool { return h.Cards[i].Value < h.Cards[j].Value }

func (t *table) Len() int           { return len(t.Hands) }
func (t *table) Swap(i, j int)      { t.Hands[i], t.Hands[j] = t.Hands[j], t.Hands[i] }
func (t *table) Less(i, j int) bool { return t.Hands[i].Value < t.Hands[j].Value }

func buildDeck() deck {
	var d = deck{}
	for _, v := range suits {
		for i := 1; i < 14; i++ {
			var a = i
			if i == 1 {
				a = 14
			}
			c := card{Low: i, Suit: v, High: i, Value: a}
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
