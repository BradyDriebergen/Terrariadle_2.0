package db

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

// Connect initializes a MongoDB client and stores it in db.Client
func Connect() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MongoDB URI not defined")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	// Verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping failed:", err)
	}

	log.Println("✅ Connected to MongoDB Atlas")
	Client = client
}

// GetCollection is a helper to grab a collection
func GetCollection(database, collection string) *mongo.Collection {
	return Client.Database(database).Collection(collection)
}
