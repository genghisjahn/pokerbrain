package poker

import "fmt"

type Card struct {
	Low  int    `json:"-"`
	High int    `json:"-"`
	Suit string `json:"suit"`
	Name string `json:"Name"`
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Name, c.Suit)
}
