package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var quote = NewQuote()

func randomAPI(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(quote.random())
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


func uptimerobot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello uptimerobot")
}