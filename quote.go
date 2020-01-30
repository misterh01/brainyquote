package main

import (
	"go.mongodb.org/mongo-driver/bson"
)

// Quote = single quote info
type Quote struct {
	ID			string	`bson:"_id,omitempty"`
	Topic		string	`bson:"topic,omitempty"`
	Quote		string	`bson:"quote,omitempty"`
	Author	string	`bson:"author,omitempty"`
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