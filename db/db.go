package db

import (
	"context"
	"log"
	"os"
	"time" // Import time for the context timeout

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func ConnectToMongo() (*mongo.Client, error) {
	// MongoDB connection string
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Getting username and password from .env
	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")

	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})

	// Set up a context with a timeout for the connection attempt
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Ping the database to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not ping database: ", err)
		return nil, err
	}

	log.Println("Successfully connected and pinged MongoDB!")

	return client, nil
}
