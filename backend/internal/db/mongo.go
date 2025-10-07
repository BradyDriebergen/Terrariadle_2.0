package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"terrariadle-backend/internal/types"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Connect initializes a MongoDB client and stores it in db.Client
func Connect() {
	// paths := []string{".env", "../.env", "../../.env"}
	err := godotenv.Load("../.env")
	if err != nil {
		wd, _ := os.Getwd()
		log.Printf("dotenv load failed: %v (cwd=%s). Falling back to system env", err, wd)
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MongoDB URI not defined")
	}

	// Use a context with a timeout to prevent the application from hanging.
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

// DisconnectDB closes the MongoDB connection. It's good practice to defer this call.
func DisconnectDB() {
	if Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := Client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
		fmt.Println("Disconnected from MongoDB.")
	}
}

// GetCollection is a helper to grab a collection
func GetCollection(database, collection string) *mongo.Collection {
	return Client.Database(database).Collection(collection)
}

func GetGuessRecord(collection *mongo.Collection, filter any) (*GuessDocument, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user GuessDocument
	res := collection.FindOne(ctx, filter)
	err := res.Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to find record %v", err)
	}

	return &user, nil
}

func GetGameData(collection *mongo.Collection, filter any) (*types.GameData, error) {
	if collection.Name() != "daily_data" {
		return &types.GameData{}, fmt.Errorf("tried to get game data from invalid collection: %s", collection.Name())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var data types.GameData
	res := collection.FindOne(ctx, filter)
	err := res.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to find record %v", err)
	}

	return &data, nil
}

// Inserts a record into a collection
func InsertRecord(collection *mongo.Collection, data any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to insert record: %v", err)
	}

	fmt.Printf("Inserted a single document with ID: %v\n", res.InsertedID)
	return nil
}

// UpdateRecord updates a single record in a collection.
func UpdateRecord(collection *mongo.Collection, filter, update any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update record: %v", err)
	}

	fmt.Printf("Matched %v document(s) and updated %v document(s).\n", res.MatchedCount, res.ModifiedCount)
	return nil
}

func UpsertRecord(collection *mongo.Collection, filter, update any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	res, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return fmt.Errorf("failed to upsert record: %v", err)
	}

	fmt.Printf("Matched %v, modified %v, upsertedID %v\n", res.MatchedCount, res.ModifiedCount, res.UpsertedID)
	return nil
}

// Delete record from a collection
func DeleteRecord(collection *mongo.Collection, filter any) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to delete record: %v", err)
	}

	fmt.Printf("Deleted %v document\n", res.DeletedCount)
	return nil
}
