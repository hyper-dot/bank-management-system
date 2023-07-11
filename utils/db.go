package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func CreateConnection() *mongo.Client {
	ctx := context.TODO()
	URI := "mongodb://127.0.0.1:27017"
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
func CreateAccountCollection(client *mongo.Client) *mongo.Collection {
	collection := client.Database("Bank").Collection("accounts")
	return collection
}

var client = CreateConnection()
var Collection = CreateAccountCollection(client)
