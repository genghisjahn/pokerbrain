package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/genghisjahn/pokerbrain/poker"
)

func main() {
	log.Println("Started")
	http.HandleFunc("/", handscoreHandler)
	http.HandleFunc("/players/score", playersScoreHandler)
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

/* Players hands json

[
   {
      "Name":"Bill",
      "Hand":[
         {
            "suit":"♧",
            "Name":"3"
         },
         {
            "suit":"♧",
            "Name":"6"
         },
         {
            "suit":"♧",
            "Name":"7"
         },
         {
            "suit":"♧",
            "Name":"8"
         },
         {
            "suit":"♧",
            "Name":"A"
         }
      ]
   },
   {
      "Name":"Frank",
      "Hand":[
         {
            "suit":"♧",
            "Name":"6"
         },
         {
            "suit":"♤",
            "Name":"6"
         },
         {
            "suit":"♢",
            "Name":"10"
         },
         {
            "suit":"♧",
            "Name":"A"
         },
         {
            "suit":"♤",
            "Name":"A"
         }
      ]
   },
   {
      "Name":"David",
      "Hand":[
         {
            "suit":"♤",
            "Name":"7"
         },
         {
            "suit":"♢",
            "Name":"7"
         },
         {
            "suit":"♢",
            "Name":"10"
         },
         {
            "suit":"♡",
            "Name":"10"
         },
         {
            "suit":"♧",
            "Name":"A"
         }
      ]
   },
   {
      "Name":"Edward",
      "Hand":[
         {
            "suit":"♤",
            "Name":"7"
         },
         {
            "suit":"♡",
            "Name":"7"
         },
         {
            "suit":"♧",
            "Name":"8"
         },
         {
            "suit":"♤",
            "Name":"8"
         },
         {
            "suit":"♧",
            "Name":"A"
         }
      ]
   },
   {
      "Name":"Jon",
      "Hand":[
         {
            "suit":"♤",
            "Name":"7"
         },
         {
            "suit":"♧",
            "Name":"8"
         },
         {
            "suit":"♢",
            "Name":"10"
         },
         {
            "suit":"♧",
            "Name":"A"
         },
         {
            "suit":"♡",
            "Name":"A"
         }
      ]
   },
   {
      "Name":"Ivan",
      "Hand":[
         {
            "suit":"♤",
            "Name":"7"
         },
         {
            "suit":"♧",
            "Name":"8"
         },
         {
            "suit":"♢",
            "Name":"10"
         },
         {
            "suit":"♤",
            "Name":"10"
         },
         {
            "suit":"♧",
            "Name":"A"
         }
      ]
   },
   {
      "Name":"Greg",
      "Hand":[
         {
            "suit":"♧",
            "Name":"8"
         },
         {
            "suit":"♡",
            "Name":"8"
         },
         {
            "suit":"♢",
            "Name":"10"
         },
         {
            "suit":"♤",
            "Name":"J"
         },
         {
            "suit":"♧",
            "Name":"A"
         }
      ]
   },
   {
      "Name":"Charles",
      "Hand":[
         {
            "suit":"♧",
            "Name":"8"
         },
         {
            "suit":"♢",
            "Name":"10"
         },
         {
            "suit":"♤",
            "Name":"Q"
         },
         {
            "suit":"♢",
            "Name":"K"
         },
         {
            "suit":"♧",
            "Name":"A"
         }
      ]
   },
   {
      "Name":"Henry",
      "Hand":[
         {
            "suit":"♤",
            "Name":"7"
         },
         {
            "suit":"♧",
            "Name":"8"
         },
         {
            "suit":"♢",
            "Name":"10"
         },
         {
            "suit":"♧",
            "Name":"K"
         },
         {
            "suit":"♧",
            "Name":"A"
         }
      ]
   },
   {
      "Name":"Adam",
      "Hand":[
         {
            "suit":"♤",
            "Name":"7"
         },
         {
            "suit":"♧",
            "Name":"8"
         },
         {
            "suit":"♢",
            "Name":"10"
         },
         {
            "suit":"♢",
            "Name":"J"
         },
         {
            "suit":"♧",
            "Name":"A"
         }
      ]
   }
]

*/
