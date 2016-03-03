package poker

import "fmt"

type deck struct {
	Cards []Card
}

func (d *deck) Deal() Card {
	card := d.Cards[0]
	d.Cards = append(d.Cards[:0], d.Cards[1:]...)
	return card
}

func buildDeck() Deck {
	var d = Deck{}
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
			d.Cards = append(d.Cards, c)
		}
	}
	return d
}
