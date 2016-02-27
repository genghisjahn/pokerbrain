package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

var suits = []string{"♤", "♡", "♢", "♧"}

const (
	HighCard      = 100
	Pair          = 200
	TwoPair       = 300
	ThreeOfKind   = 400
	Straight      = 500
	Flush         = 600
	FullHouse     = 700
	FourofKind    = 800
	StraightFlush = 900
	LOW           = false
	HIGH          = true
)

type card struct {
	Low  int
	High int
	Suit string
	Name string
}

type deck struct {
	Cards []card
}

type player struct {
	Cards [2]card
}

type table struct {
	Cards [5]card
}

type hand struct {
	Cards [5]card
}

func main() {
	deck := buildDeck()
	deck.Shuffle(5)
	for h := 0; h < 10000; h++ {
		hd1 := hand{}
		hd2 := hand{}
		deck = buildDeck()
		deck.Shuffle(5)
		for k := range hd1.Cards {
			hd1.Cards[k] = deck.Deal()
			hd2.Cards[k] = deck.Deal()
		}
		sort.Sort(hd1)
		sort.Sort(hd2)
		fmt.Println(hd1.Score(), hd1.Cards)
		fmt.Println(hd2.Score(), hd2.Cards)
		fmt.Println("-----")
		if hd1.Score() == HighCard || hd2.Score() == HighCard {
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
	suit := h.Cards[0].Suit
	flush = true
	for _, c := range h.Cards {
		if c.Suit != suit {
			flush = false
		}
	}
	unique = checkunique(h.Cards)
	straight = checkstraight(h.Cards)
	ranks := checkranks(h.Cards)
	if straight && flush {
		return StraightFlush
	}
	if ranks == FullHouse {
		return FullHouse
	}
	if flush {
		return Flush
	}
	if straight {
		return Straight
	}
	if ranks > 0 {
		return ranks
	}

	if !straight && !flush && unique {
		return HighCard
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
	c := make(map[int]int)
	for _, v := range cards {
		if _, ok := c[v.High]; ok {
			return false
		}
		c[v.High] = v.High
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

func compareHands(hands ...hand) []hand {
	winners := []hand{}
	for _, h := range hands {
		sort.Sort(h)
	}
	return winners
}

func (a hand) Len() int           { return len(a.Cards) }
func (a hand) Swap(i, j int)      { a.Cards[i], a.Cards[j] = a.Cards[j], a.Cards[i] }
func (a hand) Less(i, j int) bool { return a.Cards[i].High < a.Cards[j].High }

func buildDeck() deck {
	var d = deck{}
	for _, v := range suits {
		for i := 1; i < 14; i++ {
			c := card{Low: i, Suit: v, High: i}
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
