package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/genghisjahn/pokerbrain/poker"
)

func main() {
	log.Println("Started")
	http.HandleFunc("/hand/score", handscoreHandler)
	http.HandleFunc("/players/score", handScoreAll)
	http.HandleFunc("/hand/best", handBestHandler)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func playersScoreHandler(w http.ResponseWriter, r *http.Request) {
	//Build an anonymous struct here

	players := []struct {
		Name string
		Hand []poker.Card `json:"Cards"`
	}{}

	method := r.Method
	if method != "POST" {
		http.Error(w, fmt.Sprintf("%s not allowed", method), http.StatusMethodNotAllowed)
		return
	}

	poker.BuildDeck()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&players)
	if err != nil {
		panic(err)
	}

	t := poker.Table{}
	for kp, p := range players {
		tp := poker.Player{}
		tp.Name = p.Name
		for k, v := range p.Hand {
			val, valErr := strconv.Atoi(p.Name)
			if valErr != nil {
				if v.Name == "A" {
					v.High = 14
					v.Low = 1
				}
				if v.Name == "K" {
					v.High = 13
					v.Low = 13
				}
				if v.Name == "Q" {
					v.High = 12
					v.Low = 12
				}
				if v.Name == "J" {
					v.High = 11
					v.Low = 11
				}
			} else {
				v.High = val
				v.Low = val
			}
			tp.Hand.Cards[k] = v
			fmt.Println("card ", v, v.Low, v.High)
		}
		tp.Hand.SetScore()
		fmt.Println(kp+1, tp.Hand, tp.Hand.Score, tp.Name)

		t.Players = append(t.Players, tp)
	}
	pplayers := t.SortPlayerHands()
	for _, i := range pplayers {
		fmt.Println(i)
	}
}

func handScoreAll(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != "POST" {
		http.Error(w, fmt.Sprintf("%s not allowed", method), http.StatusMethodNotAllowed)
		return
	}

	playerScore := poker.ScoreAll{}

	poker.BuildDeck2Char()
	cardDupes := make(map[string]string)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&playerScore)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error decoding json: %v", err), http.StatusBadRequest)
		return
	}
	cardCheck := []poker.Card{}
	playerDupes := make(map[string]string)
	for _, p := range playerScore.Players {
		if _, ok := playerDupes[p.Name]; ok {
			http.Error(w, fmt.Sprintf("The player %v exists more than once ", p), http.StatusBadRequest)
			return
		}
		cardCheck = append(cardCheck, p.Pocket[0], p.Pocket[1])
		playerDupes[p.Name] = p.Name
	}
	for _, c := range playerScore.Community {
		cardCheck = append(cardCheck, c)
	}
	for _, v := range cardCheck {
		if _, ok := cardDupes[v.String()]; ok {
			http.Error(w, fmt.Sprintf("The card %v exists more than once in the cards supplied", v), http.StatusBadRequest)
			return
		}
		cardDupes[v.String()] = v.String()
	}
	numPlayers := len(playerScore.Players)
	if numPlayers < 2 || numPlayers > 10 {
		http.Error(w, fmt.Sprintf("Number of players must be between 2 and 10.  You supplied: %v", numPlayers), http.StatusBadRequest)
	}

	for k, p := range playerScore.Players {
		cards := []poker.Card{}
		vp1 := poker.DeckCardMapChar2[p.Pocket[0].String2()]
		vp2 := poker.DeckCardMapChar2[p.Pocket[1].String2()]
		cards = append(cards, vp1, vp2)
		for _, cm := range playerScore.Community {
			cmv := poker.DeckCardMapChar2[cm.String2()]
			cards = append(cards, cmv)
		}
		result, err := poker.GetBestHand(cards)
		if err != nil {
			http.Error(w, fmt.Sprintf("Something went wrong: %v\n", err), http.StatusServiceUnavailable)
			return
		}
		p.BestHand = result.Name
		p.Score = result.Score
		p.BestCards = result.Best5
		playerScore.Players[k] = p
		fmt.Println("Player:", p)
	}
	sort.Sort(sort.Reverse(&playerScore))
	jdata, _ := json.Marshal(playerScore)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jdata)
}

func handBestHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != "GET" {
		http.Error(w, fmt.Sprintf("%s not allowed", method), http.StatusMethodNotAllowed)
		return
	}
	poker.BuildDeck2Char()
	dupes := make(map[string]string)
	vals := strings.Split(r.FormValue("h"), "|")
	cardlen := len(vals)
	if cardlen < 5 || cardlen > 7 {
		http.Error(w, fmt.Sprintf("You need between 5 and 7 cards, you supplied %d\n", cardlen), http.StatusBadRequest)
		return
	}
	cards := make([]poker.Card, cardlen)
	for k, v := range vals {
		lowV := strings.ToLower(v)
		if strings.HasSuffix(lowV, "1") {
			tempS := string(lowV[0])
			lowV = tempS + "a"
		}
		if _, ok := dupes[lowV]; ok {
			http.Error(w, fmt.Sprintf("The card %v exists more than once in the hand", lowV), http.StatusBadRequest)
			return
		}
		dupes[lowV] = lowV
		cards[k] = poker.DeckCardMapChar2[lowV]
	}
	result, err := poker.GetBestHand(cards)
	if err != nil {
		http.Error(w, fmt.Sprintf("Something went wrong: %v\n", err), http.StatusServiceUnavailable)
		return
	}
	jdata, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jdata)
}

func handscoreHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != "GET" {
		http.Error(w, fmt.Sprintf("%s not allowed", method), http.StatusMethodNotAllowed)
		return
	}
	poker.BuildDeck()
	hand := poker.Hand{}
	vals := strings.Split(r.FormValue("h"), "|")
	dupes := make(map[string]string)
	for k, v := range vals[:5] {
		if _, ok := dupes[v]; ok {
			http.Error(w, fmt.Sprintf("The card %v exists more than once in the hand", v), http.StatusBadRequest)
			return
		}
		dupes[v] = v
		hand.Cards[k] = poker.DeckCardMap[v]
	}
	hand.SetScore()
	jdata, _ := json.Marshal(hand)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jdata)
}
