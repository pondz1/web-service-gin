package database

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

const (
	CONNECTIONSTRING = "mongodb+srv://dbUser:u99OTM89jo2fEWb8@dbdemo.jsemy.mongodb.net/test?retryWrites=true&w=majority"
	DB               = "demo"
	USER             = "users"
	PRODUCT          = "products"
	REWARD           = "rewards"
	RECEIPT          = "receipts"
)

func ConnectDatabase() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		// Connect to MongoDB
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(ctx, nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError

}
