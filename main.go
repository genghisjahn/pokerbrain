package main

import (
	"fmt"
	"math/rand"
	"time"
)

var suits = []string{"Spades", "Hearts", "Clubs", "Diamonds"}

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

func main() {
	deck := buildDeck()
	deck.Shuffle(5)
	players := [10]player{}
	for k, p := range players {
		p.Cards[0] = deck.Deal()
		p.Cards[1] = deck.Deal()
		fmt.Println(k+1, p.Cards[0], p.Cards[1])
	}
	t := table{}
	for k := range t.Cards {
		t.Cards[k] = deck.Deal()
	}
	fmt.Println("Table....")
	fmt.Println("Flop:", t.Cards[0], ",", t.Cards[1], ",", t.Cards[2])
	fmt.Println("Turn:", t.Cards[3])
	fmt.Println("River:", t.Cards[4])
	fmt.Println(len(deck.Cards))
}

func (d *deck) Deal() card {
	card := d.Cards[0]
	d.Cards = append(d.Cards[:0], d.Cards[1:]...)
	return card
}

func (c card) String() string {
	return fmt.Sprintf("%s of %s", c.Name, c.Suit)
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

func buildDeck() deck {
	var d = deck{}
	for _, v := range suits {
		for i := 1; i < 14; i++ {
			c := card{Low: i, Suit: v, High: i}
			if i == 1 {
				c.High = 14
				c.Name = "Ace"
			}
			if i > 1 && i < 11 {
				c.Name = fmt.Sprintf("%v", i)
			}

			if i == 11 {
				c.Name = "Jack"
			}
			if i == 12 {
				c.Name = "Queen"
			}
			if i == 13 {
				c.Name = "King"
			}
			d.Cards = append(d.Cards, c)
		}
	}
	return d
}
