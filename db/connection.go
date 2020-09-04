package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// ConnectDb : Database connexion method
func ConnectDb() *mongo.Database {
	client := InitiateMongoClient()
	return client.Database("todo")
}

// ConnectGridFs : Database connexion method to GridFS
func ConnectGridFs() *mongo.Database {
	client := InitiateMongoClient()
	return client.Database("dmpfiles")
}

// InitiateMongoClient : Database connexion method
func InitiateMongoClient() *mongo.Client {
	var err error
	var client *mongo.Client
	uri := "mongodb://localhost:27017/"
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		log.Fatal(err)
	}
	return client
}
