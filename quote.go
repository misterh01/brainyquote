package main

import (
	"strings"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/gocolly/colly"
)

// Quote = single quote info
type Quote struct {
	ID			string	`bson:"_id,omitempty"`
	Topic		string	`bson:"topic,omitempty"`
	Quote		string	`bson:"quote,omitempty"`
	Author	string	`bson:"author,omitempty"`
}

// DailyQuote quote of the day
type DailyQuote struct {
	Quote		string
	Author	string
	Imgurl	string
}

// NewQuote return new Quote
func NewQuote() *Quote {
	return &Quote{}
}

func (q *Quote) random() Quote {
	ag := bson.D{
		{Key: "$sample", Value: bson.D{
			{Key: "size", Value: 1},
		}},
	}

	return find(ag)[0]
}

func (q *Quote) topicByName(name string) Quote {
	a := bson.D{
		{Key: "$match", Value: bson.D{{Key:"topic", Value:name}}},
	}

	d := find(a)
	if len(d) > 0 {
		return randomElem(d)
	}
	return Quote{}
}

func (q *Quote) quoteByID(id string) Quote {
	a := bson.D{
		{Key: "$match", Value: bson.D{{Key:"_id", Value:id}}},
	}
	d := find(a)

	if len(d) > 0 {
		return d[0]
	}
	return Quote{}
}

func (q *Quote) quoteOfTheDay() DailyQuote {
	var quote DailyQuote
	
	c := colly.NewCollector()
	c.OnHTML(".bqTopLevel .qotd_days .container .bqQt .oncl_q .img-responsive", func(e *colly.HTMLElement) {
		alt := strings.Split(e.Attr("alt"), " - ")
		quote.Quote = alt[0]
		quote.Author = alt[1]
    quote.Imgurl = "https://www.brainyquote.com" + e.Attr("data-img-url")
	})
	c.Visit("https://www.brainyquote.com/quote_of_the_day")

	return quote
}
