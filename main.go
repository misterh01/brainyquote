package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"math/rand"
	"time"

	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MongoDBInit()

	rand.Seed(time.Now().UnixNano())
	r := mux.NewRouter()

	r.HandleFunc("/", randomAPI)
	r.HandleFunc("/findbyid", findByIDAPI)
	r.HandleFunc("/topic/{topic}", topicNameAPI)
	r.HandleFunc("/daily", dailyQuoteAPI)

	fmt.Println("Server Started!")
	log.Fatal(http.ListenAndServe(GetPort(), r))
}


// GetPort - heroku port or default = 8080
func GetPort() string {
  p := os.Getenv("PORT")
  if p != "" {
    return ":" + p
  }
  return ":8080"
}
