package poker

import "strconv"

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

func checkunique(cards [5]Card) bool {
	c := make(map[int]int)
	for _, v := range cards {
		if _, ok := c[v.High]; ok {
			return false
		}
		c[v.High] = v.High
	}
	return true
}

func checkstraight(cards [5]Card) bool {
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

func checkranks(cards [5]Card) (string, [2]int) {
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
		var firstset = false
		for k, c1 := range c {
			if c1 == 2 {
				if !firstset {
					vals[0] = k
					firstset = true
				} else {
					vals[1] = k
				}
			}
			// if c1 == 2 && vals[0] == 0 {
			// 	vals[0] = k
			// }
			// if c1 == 2 && vals[1] == 0 {
			// 	vals[1] = k
			// }
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

func checkMaxDiff(min, max int) bool {
	if max-min == 5 {
		return true
	}
	return false
}

func getbounds(highlow bool, c [5]Card) (int, int) {
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

func GetCardCombinations(cards []Card) []Hand {
	var hands = []Hand{}
	length := len(cards)
	for a := 0; a < length-4; a++ {
		for b := (a + 1); b < length-3; b++ {
			for c := (b + 1); c < length-2; c++ {
				for d := (c + 1); d < length-1; d++ {
					for e := (d + 1); e < length; e++ {
						h := Hand{}
						h.Cards[0] = cards[a]
						h.Cards[1] = cards[b]
						h.Cards[2] = cards[c]
						h.Cards[3] = cards[d]
						h.Cards[4] = cards[e]
						hands = append(hands, h)
					}
				}
			}
		}
	}
	return hands
}
