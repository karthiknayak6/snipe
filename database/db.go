package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type Counter struct {
	ID   string `bson:"_id"`
	Seq  int    `bson:"seq"`
}

var Client *mongo.Client

var Db *mongo.Database

func GetNextSequence(client *mongo.Client, collectionName string) (int, error) {
	collection := client.Database("snipe").Collection("counters")

	filter := bson.M{"_id": collectionName}
	update := bson.M{"$inc": bson.M{"seq": 1}}


	opt := options.FindOneAndUpdate().
	SetReturnDocument(options.After).  // Return the updated document
	SetUpsert(true)                    // Perform upsert if the document doesn't exist


	var result Counter	
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, opt).Decode(&result)
	if err != nil {
		return 0, err
	}

	return result.Seq, nil
}



func CreateConnection() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	Db = Client.Database("snipe")
	
	if err != nil {
		log.Fatal(err)
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func TerminateConnection() {
	err := Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
