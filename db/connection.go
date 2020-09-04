package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDb : Database connexion method
func ConnectDb() *mongo.Database {
	ctx := context.Background()
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://localhost:27017/todo"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database("todo")
}
