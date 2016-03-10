package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/genghisjahn/pokerbrain/poker"
)

func main() {
	log.Println("Started")
	http.HandleFunc("/hand/score", handscoreHandler)
	http.HandleFunc("/players/score", playersScoreHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func playersScoreHandler(w http.ResponseWriter, r *http.Request) {
	//Build an anonymous struct here
	method := r.Method
	if method != "POST" {
		http.Error(w, fmt.Sprintf("%s not allowed", method), http.StatusMethodNotAllowed)
		return
	}
	poker.BuildDeck()
	decoder := json.NewDecoder(r.Body)
	var p []poker.Player
	err := decoder.Decode(&p)
	if err != nil {
		panic(err)
	}
	log.Println("***", p[0].Hand)
	for k := range p {
		_ = k
		//p[k].Hand = p.Cards
	}
	t := poker.Table{}
	p = t.SortPlayerHands()
	for _, i := range p {
		fmt.Println(i)
	}
	log.Println(p)
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
	for k, v := range vals[:5] {
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
