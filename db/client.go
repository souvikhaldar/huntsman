package db

import (
	"context"
	"fmt"
	"huntsman/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var mongoIPCollection *mongo.Collection

func initializeMongoDB(con config.Config) error {
	// create the mongodb client
	clientOptions := options.Client().ApplyURI(con.MongoURI)
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Failed to create mongo db connection: ", err)
		return err
	}
	if err := mongoClient.Ping(context.TODO(), nil); err != nil {
		fmt.Println("Failed to ping mongo db instance: ", err)
		return err
	}
	mongoIPCollection = mongoClient.Database(con.MongoDatabase).Collection(con.MongoCollection)
	return nil
}
