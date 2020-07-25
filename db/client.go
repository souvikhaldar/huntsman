package db

import (
	"context"
	"fmt"
	"huntsman/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoIPCollection *mongo.Collection

func InitializeMongoDB(con config.Config) error {
	// create the mongodb client
	clientOptions := options.Client().ApplyURI(con.MongoURI)
	MongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Failed to create mongo db connection: ", err)
		return err
	}
	if err := MongoClient.Ping(context.TODO(), nil); err != nil {
		fmt.Println("Failed to ping mongo db instance: ", err)
		return err
	}
	MongoIPCollection = MongoClient.Database(con.MongoDatabase).Collection(con.MongoCollection)
	return nil
}
