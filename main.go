package main

import (
	"fmt"

	"github.com/genghisjahn/pokerbrain/poker"
)

func main() {
	deck := poker.BuildDeck()
	deck.Shuffle(5)
	hands := make([]poker.Hand, 10, 10)

	for i, h := range hands {
		for k := range h.Cards {
			hands[i].Cards[k] = deck.Deal()
		}
	}

	t := poker.Table{}
	t.Hands = append(t.Hands, hands...)

	ch := t.CompareHands()
	fmt.Println("Compare:")
	for _, c := range ch {
		fmt.Println(c)
	}

}
