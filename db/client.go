package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoIPCollection *mongo.Collection

func InitializeMongoDB() {
	// create the mongodb client
	clientOptions := options.Client().ApplyURI(
		"mongodb+srv://souvikhaldar:Pl%40y1tMongo@cluster0.fe2ea.mongodb.net/huntsman?retryWrites=true&w=majority",
	)
	MongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Failed to create mongo db connection: ", err)
		return
	}
	if err := MongoClient.Ping(context.TODO(), nil); err != nil {
		fmt.Println("Failed to ping mongo db instance: ", err)
		return
	}
	MongoIPCollection = MongoClient.Database("huntsman").Collection("ipdata")
}
