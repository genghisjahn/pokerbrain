package poker

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var DeckCardMap map[string]Card
var DeckCardMapChar2 map[string]Card

var SuitMap map[string]string

type Deck struct {
	Cards []Card
}

func (d *Deck) Deal() Card {
	card := d.Cards[0]
	d.Cards = append(d.Cards[:0], d.Cards[1:]...)
	return card
}

func (d *Deck) Shuffle(num int) {
	rand.Seed(time.Now().UnixNano())
	for n := 0; n < num; n++ {
		for i := range d.Cards {
			j := rand.Intn(i + 1)
			d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
		}
	}
}

func BuildDeck() Deck {
	var d = Deck{}
	SuitMap = make(map[string]string)
	SuitMap["s"] = "s"
	SuitMap["h"] = "h"
	SuitMap["d"] = "d"
	SuitMap["c"] = "c"
	DeckCardMap = make(map[string]Card)
	for _, v := range suits {
		for i := 1; i < 14; i++ {
			c := Card{Low: i, Suit: v, High: i}
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
			DeckCardMap[fmt.Sprintf("%s%v", SuitMap[c.Suit], c.High)] = c
			d.Cards = append(d.Cards, c)
		}
	}
	return d
}

func BuildDeck2Char() Deck {
	var d = Deck{}
	SuitMap = make(map[string]string)
	SuitMap["s"] = "s"
	SuitMap["h"] = "h"
	SuitMap["d"] = "d"
	SuitMap["c"] = "c"
	DeckCardMapChar2 = make(map[string]Card)
	for _, v := range suits {
		for i := 1; i < 14; i++ {
			c := Card{Low: i, Suit: v, High: i}
			if i == 1 {
				c.High = 14
				c.Name = "A"
			}
			if i > 1 && i < 10 {
				c.Name = fmt.Sprintf("%v", i)
			}
			if i == 10 {
				c.Name = "T"
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
			DeckCardMapChar2[fmt.Sprintf("%s%v", SuitMap[c.Suit], strings.ToLower(c.Name))] = c
			d.Cards = append(d.Cards, c)
		}
	}
	return d
}
