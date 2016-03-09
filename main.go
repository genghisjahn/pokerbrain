package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/genghisjahn/pokerbrain/poker"
)

func main() {
	log.Println("Started")
	http.HandleFunc("/score", scoreHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func scoreHandler(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/hand?h=h14|s14|s10|c8|d4
	poker.BuildDeck()
	hand := poker.Hand{}
	vals := strings.Split(r.FormValue("h"), "|")
	for k, v := range vals[:5] {
		hand.Cards[k] = poker.DeckCardMap[v]
	}
	hand.SetScore()
	jdata, _ := json.Marshal(hand)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jdata)
}
