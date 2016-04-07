package poker

type ScorePlayer struct {
	Name     string  `json:"name"`
	Pocket   [2]Card `json:"pocket"`
	Score    string  `json:"score,omitempty"`
	BestHand string  `json:"best_hand,omitempty"`
}

type ScoreAll struct {
	Players   []ScorePlayer `json:"players"`
	Community []Card        `json:"community"`
}
