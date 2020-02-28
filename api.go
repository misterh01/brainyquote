package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var quote = NewQuote()

func randomAPI(w http.ResponseWriter, r *http.Request) {
	total, err := strconv.Atoi(queryValue(r, "total"))
	if err != nil || total <= 0 || total > 10{
		total = 1
	}

	j, _ := json.Marshal(quote.random(total))
	fmt.Fprintln(w, string(j))
}

func topicNameAPI(w http.ResponseWriter, r *http.Request) {
	n := mux.Vars(r)["topic"]
	j, _ := json.Marshal(quote.topicByName(n))
	fmt.Fprintln(w, string(j))
}


func findByIDAPI(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(quote.quoteByID(queryValue(r, "id")))
	fmt.Fprintln(w, string(j))
}

func dailyQuoteAPI(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(quote.quoteOfTheDay())
	fmt.Fprintln(w, string(j))
}
