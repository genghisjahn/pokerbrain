package poker

type ScorePlayer struct {
	Name      string  `json:"name"`
	Pocket    [2]Card `json:"pocket"`
	Score     string  `json:"score,omitempty"`
	BestHand  string  `json:"best_hand,omitempty"`
	BestCards [5]Card `json:"best_cards,omitempty"`
}

func (sa *ScoreAll) Len() int           { return len(sa.Players) }
func (sa *ScoreAll) Swap(i, j int)      { sa.Players[i], sa.Players[j] = sa.Players[j], sa.Players[i] }
func (sa *ScoreAll) Less(i, j int) bool { return sa.Players[i].Score < sa.Players[j].Score }

type ScoreAll struct {
	Players   []ScorePlayer `json:"players"`
	Community []Card        `json:"community"`
}
