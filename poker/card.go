package poker

import "fmt"

type Card struct {
	Low  int
	High int
	Suit string
	Name string
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Name, c.Suit)
}
