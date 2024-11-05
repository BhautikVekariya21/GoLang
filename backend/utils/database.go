package utils

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// ConnectDB initializes the MongoDB client
func ConnectDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get MongoDB URI from environment variable
	mongoURI := os.Getenv("MONGODB_URI")

	// Create a new MongoDB client
	ClientOptions := options.Client().ApplyURI(mongoURI)
	Client, err = mongo.Connect(context.TODO(), ClientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Optionally, ping the database to check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = Client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
}
