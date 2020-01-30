package main

import (
	// "fmt"
	"os"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	dbName = "db"
	dbColl = "q_collection"
)


// Client mongodb client
var Client *mongo.Client


// MongoDBInit init mongodb
func MongoDBInit() {
	var err error
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	Client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
}


func find(ag bson.D) []Quote {
	collection := Client.Database(dbName).Collection(dbColl)
	cur, _ := collection.Aggregate(context.TODO(), mongo.Pipeline{ag})
	var result []Quote

	for cur.Next(context.TODO()) {
		var elem Quote
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal("Something Went Wrong!")
		}
		result = append(result, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal("Something Went Wrong!")
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return result
}