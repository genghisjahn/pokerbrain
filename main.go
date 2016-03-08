package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	log.Println("Started")
	http.HandleFunc("/score", scoreHandler)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func scoreHandler(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/hand?h=h14|s14|s10|c8|d4
	vals := strings.Split(r.FormValue("h"), "|")
	for _, v := range vals {
		fmt.Println(v)
	}
	http.Error(w, "Not Implemented.", http.StatusInternalServerError)
}
