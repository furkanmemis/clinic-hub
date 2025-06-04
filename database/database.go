package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection(databaseName string, collectionName string) *mongo.Collection {
	var collection *mongo.Collection

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://clinic-hub-mongodb-1:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Connection err: ", err)
	}

	collection = client.Database(databaseName).Collection(collectionName)
	return collection
}
